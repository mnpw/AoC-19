fuel = 0

with open("input", "r") as f:
    for i in f:
        a = int(i)
        while a > 0:
            a = a//3 - 2
            fuel += max(0, a)
            
print(fuel)
