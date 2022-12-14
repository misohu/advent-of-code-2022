from collections import namedtuple

point = namedtuple("point", "x y")

FILE_NAME = "input.txt"

def read_file(file_name):
    res = []
    with open(file_name, "r") as f:
        for line in f:
            parts = line.strip().split("->")
            tmp = [part.strip().split(",") for part in parts]
            coords = [[int(x[0]), int(x[1])] for x in tmp]
            res.append(coords)
    return res

def process_walls(coords):
    memory = set()
    border = 0
    for line in coords:
        for (x1, y1), (x2, y2) in zip(line, line[1:]):
            x1, x2 = sorted([x1, x2])
            y1, y2 = sorted([y1, y2])
            for x in range(x1, x2 + 1):
                for y in range(y1, y2 + 1):
                    memory.add(point(x, y))
                    if y + 1 > border:
                        border = y + 1
    return border, memory

def fall(border, memory, start):
    counter = 0
    while point(start, 0) not in memory:
        s = point(start, 0)
        while True:
            if s.y >= border:
                break
            if point(s.x, s.y + 1) not in memory:
                s = point(s.x, s.y + 1)
                continue
            if point(s.x - 1, s.y + 1) not in memory:
                s = point(s.x - 1, s.y + 1)
                continue
            if point(s.x + 1, s.y + 1) not in memory:
                s = point(s.x + 1, s.y + 1)
                continue
            break
        memory.add(s)
        counter += 1
    return counter

if __name__ == "__main__":
   coords = read_file(FILE_NAME)
   border, memory = process_walls(coords)
   print(fall(border, memory, 500))