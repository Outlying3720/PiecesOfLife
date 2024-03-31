import glob
import string

files = glob.glob("/workspace/pla/output/workspace/pla/tools/cfgs/scannet_models/spconv_clip_adamw_oneformer_base10_caption/oneformer_b10_b32_cap_wbinary_woaux_wonormalize_scratch2/ckpt/*.pth")

if len(files) > 10:
    print('Step delete.')
    filenums = []
    for file in files:
        filenum = file.split('_')[-1].split('.')[0]
        filenums.append(string.atoi(filenum))
    sort(filenums)
    print(filenums)