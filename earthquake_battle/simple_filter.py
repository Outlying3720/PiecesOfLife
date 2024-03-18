import numpy as np
import glob
import tqdm

files = glob.glob('/home/omnisky/Downloads/earthquake/test/QZ/*.txt')
# files = files[2000:]

def read_float_list_from_txt(file_path):
        float_list = []
        with open(file_path, 'r') as file:
            for line in file:
                float_list.append(float(line.strip()))
        return float_list

def get_nearby_avg(arr):
        nearby_max = 0
        nearby_avg = 0
        for i in range(len(arr) - 1):
            tmp = abs(arr[i] - arr[i + 1])
            nearby_avg += tmp
            nearby_max = max(tmp, nearby_max)
        nearby_avg /= len(arr)
        return nearby_avg, nearby_max

def process_boolean_list(bool_list):
        processed_list = bool_list.copy()  # 复制原始列表，以免修改原始列表内容

        for i in range(len(bool_list)):
            count_false = sum(1 for j in range(max(0, i-10), min(len(bool_list), i+11)) if not bool_list[j])
            if count_false > 10:  # 如果False的个数超过一半
                processed_list[i] = False
            else:
                processed_list[i] = True

        return processed_list

# for file_path in tqdm.tqdm(files):

#     data = read_float_list_from_txt(file_path)
#     # get pattern
#     # last_data = data[6000:7750] 
#     last_data = data[-5000:]

#     nearby_avg, nearby_max = get_nearby_avg(last_data)
#     print(nearby_avg, nearby_max)

#     result = []

#     for i in range(len(data) - 1):
#         tmp = abs(data[i] - data[i + 1])
#         if tmp > 4 * nearby_avg:
#             result.append(True)
#         else:
#             result.append(False)
#     result.append(False)


#     result = process_boolean_list(result)

#     result = np.array(result)
    
#     out = np.array(data)
#     out[result] = out[result] * 100
    
#     out_file = file_path.replace('/QZ/', '/result/').replace('_Q.txt', '_C.txt')
#     with open(out_file, 'w+') as f:
#         for line in out:
#             f.write(str(line))
#             f.write('\n')


# ------------------

file_path = '/home/omnisky/Downloads/earthquake/test/QZ/0a6f04e3ea2dfcb8d6752accf0099f8a4b108234_Q.txt'

data = read_float_list_from_txt(file_path)
data = np.array(data)

# def topk_elements(arr, k):
#     arr_sorted = sorted(arr, reverse=True)
#     return arr_sorted[:k], arr_sorted[-k:]

# topk, lastk = topk_elements(last_data, 10)

# print(topk)
# print('-------')
# print(lastk)


from scipy.signal import medfilt

def denoise_wave(float_list, kernel_size=3):
    denoised_wave = medfilt(float_list, kernel_size)
    return denoised_wave

denoised_wave = denoise_wave(data, kernel_size=31)

import matplotlib.pyplot as plt

def plot_wave(float_list, title, result = None):
    float_list = np.array(float_list)
    plt.figure(figsize=(10, 6))
    if result:
        x = np.array(list(range(len(float_list))))
        plt.scatter(x[result == False], float_list[result == False], label='Original Wave', color='b', marker='.')
        plt.scatter(x[result],float_list[result], label='Earthquake Wave', color='r', marker='.')
    else:
        plt.plot(float_list, label='Original Wave', color='b')
    plt.xlabel('Index')
    plt.ylabel('Value')
    plt.title(title)
    plt.legend()
    plt.grid(True)
    plt.show()

plot_wave(data, 'Noisy Wave', result=None)
plot_wave(denoised_wave, 'Denoised Wave')
