# 大乐透!
# 2024-03-01 20:44:53
# 来源于gzq给我看了他同学刮刮乐(中国红)中了500元
# 我想出了这个有保底机制的抽卡大乐透!
# 
# 原版规则
# 1. 每注中奖概率0.7%, 从第100抽开始线性增加中奖概率, 至150抽概率为100%
# 2. 中奖时有50%概率会歪, 歪了返还投注金额(后改为歪了翻50倍)
# 3. 如果上一次中奖歪了, 则下一次中奖必然不会歪
# 4. 中奖返还100倍投注金额
# 5. 从第一次投注到歪或中为止,投注金额必须保持不变
# 

import random


class DA_LE_TOU:
    def __init__(self):
        self.p = 0.07

        #result 0 lost, 1 wai, 2 zhong
        self.i = 0
        self.last_win = False
        self.last_wai = False
    def draw(self):
        if self.last_win:
            self.i = 0
        self.i += 1
        result = random.random() <= self.p + (1 - self.p) / 50 * (self.i - 100)
        
        if result == True:
            self.last_win = True
            if random.random() < 0.5 and self.last_wai == False:
                result = 1
                self.last_wai = True
            else:
                result = 2
                self.last_wai = False
        else:
            self.last_win = False
            result = 0
                
        return result

daletou = DA_LE_TOU()

money = 0
per = 1

for j in range(2500):
    result = daletou.draw()
    if result == 0:
        money -= per
    elif result == 1:
        money += 50 * per
        print(f'Draw {j} - {daletou.i}: 小保底歪,翻50倍 \t\t\t money: {money}')
    elif result == 2:
        money += 100 * per
        print(f'Draw {j} - {daletou.i}: 中了!!翻100倍 \t\t\t\t money: {money}')
        