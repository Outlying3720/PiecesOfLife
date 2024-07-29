import threading
from threading import Thread
import time

from queue import Queue

queue           = Queue(3)

class A:
    def __init__(self):
        print('###########THREADING')
        self.thread = Thread(target=A.preload_data, args=(self, ))
        
        # self.length          = len(self.data_list)
        self.Continue  = True
        self.thread.start()
        
    def _prepare_data(self, idx):
        time.sleep(1)
        return idx
    
    def preload_data(self):
        print('#############preload_data')
        global queue
        try:
            for idx in range(len(self)):
                print('put', idx)
                # print(queue)
                item = self._prepare_data(idx)
                queue.put((True , item))
        except Exception as e:              queue.put((False, e))
        finally:                            queue.put((False, StopIteration))
    
    def prepare_data(self, idx):
        global queue
        if self.Continue:
            print('before get ', idx)
            import time
            time.sleep(1)
            success, next_item = queue.get()
            print('done get ', idx)
            if success: return next_item
            else:
                self.Continue = False
                raise next_item
        else: raise StopIteration
    
    def __len__(self):
        return 100

a = A()

for i in range(100):
    print(a.prepare_data(i))
    time.sleep(5)