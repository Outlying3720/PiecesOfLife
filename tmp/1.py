    def predict(self, batch_inputs_dict, batch_data_samples):
        text_prompts = [
            data_samples.text for data_samples in batch_data_samples
        ]  # txt list

        

        # extract text feats
        tokenized = self.tokenizer.batch_encode_plus(
            text_prompts, padding='longest',
            return_tensors='pt').to(batch_inputs_dict['points'][0].device)

        if 'tokens_positive' in batch_data_samples[0]:
            tokens_positive = [
                data_samples.tokens_positive
                for data_samples in batch_data_samples
            ]
        else:
            # hack a pseudo tokens_positive during format-only inference
            tokens_positive = [[[0, 1]]
                               for _ in range(len(batch_data_samples))]
        positive_maps = self.get_positive_map(tokenized, tokens_positive)
        positive_maps = [
            positive_map.to(batch_inputs_dict['points']
                            [0].device).bool().float().unsqueeze(0)
            for positive_map in positive_maps
        ]  # each positive_map: (1, max_text_length)

        encoded_text = self.text_encoder(**tokenized)
        text_feats = self.text_feat_map(encoded_text.last_hidden_state)
        text_token_mask = tokenized.attention_mask.bool()
        text_dict = dict()
        text_dict['text_feats'] = text_feats
        text_dict['text_token_mask'] = text_token_mask  # (bs, max_text_length)
        # mind attention mask that we get from huggingface is inverse
        # because its the opposite in pytorch transformer
        # text_dict['tokenized'] = tokenized
        for i, data_samples in enumerate(batch_data_samples):
            text_token_mask = text_dict['text_token_mask'][
                i]  # (max_text_length)
            data_samples.gt_instances_3d.positive_maps = positive_maps[i]
            # (1, max_text_length)
            data_samples.gt_instances_3d.text_token_mask = \
                text_token_mask.unsqueeze(0).repeat(
                    len(positive_maps[i]), 1)

        point_feats, scores, point_xyz = self.extract_feat(
            batch_inputs_dict, batch_data_samples)
        head_inputs_dict = self.forward_transformer(point_feats, scores,
                                                    point_xyz, text_dict,
                                                    batch_data_samples)
        results_list = self.bbox_head.predict(
            **head_inputs_dict, batch_data_samples=batch_data_samples)

        for data_sample, pred_instances_3d in zip(batch_data_samples,
                                                  results_list):
            data_sample.pred_instances_3d = pred_instances_3d
        return batch_data_samples