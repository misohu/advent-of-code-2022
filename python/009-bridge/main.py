from collections import namedtuple

FILE_NAME="input.txt"

point = namedtuple("point", "x y")

def move_head(head: point, dir: str) -> point:
    if dir == "R":
        return point(head.x + 1, head.y)
    if dir == "L":
        return point(head.x - 1, head.y)
    if dir == "U":
        return point(head.x, head.y + 1)
    if dir == "D":
        return point(head.x , head.y - 1)

def is_neighbour(p1: point, p2: point) -> bool:
    return abs(p2.x - p1.x) <= 1 and abs(p2.y - p1.y) <= 1

def next_position(p1: point, p2: point) -> point:
    if is_neighbour(p1, p2):
        return p2
    
    if p1.x == p2.x:
        return point(p2.x ,(p2.y + p1.y)//2)
    
    if p1.y == p2.y:
        return point((p2.x + p1.x)//2 ,p2.y)
    
    diagonals = [[-1, -1], [1, 1], [-1, 1], [1, -1]]
    variants = list(map(lambda dia: point(p2.x + dia[0], p2.y + dia[1]), diagonals))
    for var in variants:
        if is_neighbour(p1, var):
            return var

def process_file(file_name):
    res = []
    with open(file_name, "r") as f:
        for line in f:
            dir, rep = line.split()
            for _ in range(int(rep)):
                res.append(dir)
    return res

def process_moves(head: point, tails: point, moves: list[str]): 
    visited1 = set()
    visited9 = set()

    for move in moves:
        head = move_head(head, move)
        tails[0] = next_position(head, tails[0])
        for i in range(1, len(tails)):
            tails[i] = next_position(tails[i -1], tails[i])
        visited1.add(tails[0])
        visited9.add(tails[8])
        print(head, tails[0])
    return len(visited1), len(visited9)


if __name__ == "__main__":
    input = process_file(FILE_NAME)
    print(input)
    tails = list(map(lambda _: point(0,0), range(9)))
    visited1, visited9 = process_moves(point(0,0), tails, input)
    print(visited1)
    print(visited9)