from functools import cmp_to_key

FILE_NAME = "input.txt"


def read_file(file_name):
    with open(file_name) as f:
        res = []

        while True:
            tmp = []
            tmp.append(eval(f.readline().strip()))
            tmp.append(eval(f.readline().strip()))
            new_line = f.readline()
            res.append(tmp)
            if len(new_line) == 0:
                break

        return res


def compare(left, right):
    if isinstance(left, int) and isinstance(right, int):
        return (left < right) - (left > right)
    left = [left] if isinstance(left, int) else left
    right = [right] if isinstance(right, int) else right
    for i in range(len(left)):
        if len(right) <= i:
            return -1
        d = compare(left[i], right[i])
        if d != 0:
            return d
    return len(left) < len(right)


if __name__ == "__main__":
    input = read_file(FILE_NAME)
    res = 0
    for i, pair in enumerate(input):
        x = compare(pair[0], pair[1])
        if x >= 0:
            res += i + 1
    print(res)
    second = []
    for x in input:
        second = second + x
    second.append([[2]])
    second.append([[6]])
    second.sort(key=cmp_to_key(lambda x, y: -compare(x, y)))
    print((second.index([[2]]) + 1) * (second.index([[6]]) + 1))
