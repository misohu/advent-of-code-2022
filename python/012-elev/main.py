from collections import namedtuple, deque

FILE_NAME = "input.txt"

point = namedtuple("point", "x y")

def read_file(file_name: str):
    with open(file_name, "r") as f:
        characters = [list(l.strip()) for l in f]
        res = []
        starts = []
        end = None
        for i, row in enumerate(characters):
            tmp = []
            for j, val in enumerate(row):
                match val:
                    case "S":
                        starts.append(point(i, j))
                        tmp.append(ord("a"))
                    case "a":
                        starts.append(point(i, j))
                        tmp.append(ord("a"))
                    case "E":
                        end = point(i, j)
                        tmp.append(ord("z"))
                    case _:
                        tmp.append(ord(val))
            res.append(tmp)
        return res, starts, end

def get_neighbours(p: point):
    return [
        point(p.x + off[0], p.y + off[1])
        for off in [[0, 1], [1, 0], [0, -1], [-1, 0]]
    ]

def find_path(field, starts, end):
    buffer = deque([(start, 0) for start in starts])
    seen = set()
    while buffer:
        current = buffer.popleft()
        if current[0] in seen:
            continue
        if current[0] == end:
            return current[1]
        seen.add(current[0])
        for n in get_neighbours(current[0]):
            if n.x < 0 or n.y < 0 or n.x >= len(field) or n.y >= len(field[0]):
                continue
            if field[n.x][n.y] <= field[current[0].x][current[0].y] + 1:
                buffer.append((n, current[1] + 1))
        

if __name__ == "__main__":
    field, starts, end  = read_file(FILE_NAME)
    res = find_path(field, starts, end)
    print(res)
