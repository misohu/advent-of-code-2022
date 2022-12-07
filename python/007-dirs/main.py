from collections import defaultdict
from itertools import accumulate

FILE_NAME = "input.txt"

if __name__ == "__main__":
    stack = []
    sizes = defaultdict(int)
    with open(FILE_NAME, "r") as f:
        for line in f:
            if line.startswith("$ ls") or line.startswith("dir"):
                continue
            parts  = line.split()
            if line.startswith("$ cd") and parts[2] != "..":
                stack.append(parts[2])
            elif line.startswith("$ cd") and parts[2] == "..":
                stack.pop()
            else:
                for folder in accumulate(stack, lambda x,y: x + "*" + y):
                    sizes[folder] += int(parts[0])
    
    print(sum(filter(lambda v: v < 100000, sizes.values())))
    print(min(filter(lambda v: sizes["/"] - 40000000 <= v, sizes.values())))