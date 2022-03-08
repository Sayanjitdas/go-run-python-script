import sys
import time

counter = int(sys.argv[1])
sum = 0
for i in range(counter):
    sum += i
    time.sleep(1)

print(f"COUNTER {counter} yields {sum}")