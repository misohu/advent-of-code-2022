from dataclasses import dataclass
from collections import deque

@dataclass
class Monkey:
    items: any
    transformer: any
    test: int
    if_true: int
    if_false: int


monkeys = [
    Monkey(deque([85, 77, 77]), lambda x: x*7, 19, 6,7),
    Monkey(deque([80, 99]), lambda x: x*11, 3, 3,5),
    Monkey(deque([74, 60, 74, 63, 86, 92, 80]), lambda x: x+8, 13, 0,6),
    Monkey(deque([71, 58, 93, 65, 80, 68, 54, 71]), lambda x: x+7, 7, 2,4),
    Monkey(deque([97, 56, 79, 65, 58]), lambda x: x+5, 5, 2,0),
    Monkey(deque([77]), lambda x: x+4, 11, 4,3),
    Monkey(deque([99, 90, 84, 50]), lambda x: x*x, 17, 7,1),
    Monkey(deque([50, 66, 61, 92, 64, 78]), lambda x: x+3, 2, 5,1),
]