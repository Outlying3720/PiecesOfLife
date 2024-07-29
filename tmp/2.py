    def loss(self, batch_inputs_dict: dict, batch_data_samples: SampleList,
             **kwargs) -> Union[dict, list]:
        """Calculate losses from a batch of inputs dict and data samples.

        Args:
            batch_inputs_dict (dict): The model input dict which include
                'points', 'img' keys.

                    - points (list[torch.Tensor]): Point cloud of each sample.
                    - imgs (torch.Tensor, optional): Image of each sample.

            batch_data_samples (List[:obj:`Det3DDataSample`]): The Data
                Samples. It usually includes information such as
                `gt_instance_3d`, `gt_panoptic_seg_3d` and `gt_sem_seg_3d`.

        Returns:
            dict: A dictionary of loss components.
        """
        text_prompts = [
            data_samples.text for data_samples in batch_data_samples
        ]  # txt list

        tokens_positive = [
            data_samples.tokens_positive for data_samples in batch_data_samples
        ]
        
        tokenized = self.tokenizer.batch_encode_plus(
            text_prompts, padding='longest',
            return_tensors='pt').to(batch_inputs_dict['points'][0].device)
        positive_maps = self.get_positive_map(tokenized, tokens_positive)

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
            positive_map = positive_maps[i].to(
                batch_inputs_dict['points']
                [0].device).bool().float().unsqueeze(0)  # (1, max_text_length)
            text_token_mask = text_dict['text_token_mask'][
                i]  # (max_text_length)
            data_samples.gt_instances_3d.positive_maps = positive_map
            # (1, max_text_length)
            data_samples.gt_instances_3d.text_token_mask = \
                text_token_mask.unsqueeze(0).repeat(
                    len(positive_map), 1)

        point_feats, scores, point_xyz = self.extract_feat(
            batch_inputs_dict, batch_data_samples)
        head_inputs_dict = self.forward_transformer(point_feats, scores,
                                                    point_xyz, text_dict,
                                                    batch_data_samples)
        losses = self.bbox_head.loss(**head_inputs_dict,
                                     batch_data_samples=batch_data_samples)
        return losses