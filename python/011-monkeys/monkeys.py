from dataclasses import dataclass
from collections import deque
from operator import add, mul

@dataclass
class Monkey:
    items: any
    transformer: any
    test: int
    if_true: int
    if_false: int


monkeys = [
    Monkey(deque([79, 98]), lambda x: mul(x,19), 23, 2,3),
    Monkey(deque([54, 65, 75, 74]), lambda x: add(x,6), 19, 2,0),
    Monkey(deque([79, 60, 97]), lambda x: mul(x,x), 13, 1,3),
    Monkey(deque([74]), lambda x: add(x,3), 17, 0,1),
]