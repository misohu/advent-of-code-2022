from monkeys2 import monkeys, Monkey
from math import prod

counts = {
    0:0,
    1:0,
    2:0,
    3:0,
    4:0,
    5:0,
    6:0,
    7:0
}

def round(monkeys: list[Monkey]):
    greatest_common_product = prod(monkey.test for monkey in monkeys)
    for i in range(len(monkeys)):
        monkey = monkeys[i]
        while monkey.items:
            item = monkey.items.pop()
            item %= greatest_common_product
            counts[i]+=1
            value = monkey.transformer(item)
            if value % monkey.test == 0:
                monkeys[monkey.if_true].items.append(value)
            else:
                monkeys[monkey.if_false].items.append(value)

if __name__ == "__main__":
    for i in range(10000):
        round(monkeys)
    for m in monkeys:
        print(m.items)
    print(counts)