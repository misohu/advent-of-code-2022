FILE_NAME = "input.txt"


def process_input(input):
    result = 0
    for i, row in enumerate(input):
        for j, _ in enumerate(row):
            tmp = heigh_value(i, j, input, len(input))
            result = tmp if tmp > result else result
    return result

def process_input1(input):
    result = 0
    for i, row in enumerate(input):
        for j, _ in enumerate(row):
            tmp = heigh_value1(i, j, input, len(input))
            result += 1 if tmp else 0
    return result


def heigh_value1(row, column, input, max_size):
    directions = ["r", "l", "u", "d"]
    for i in range(1, max_size):
        tmp_directions = directions.copy()
        for direction in tmp_directions:
            if direction == "r":
                if column + i == max_size:
                    return True
                elif input[row][column] <= input[row][column + i]:
                    directions.remove(direction)
            if direction == "l":
                if column - i < 0:
                    return True
                elif input[row][column] <= input[row][column - i]:
                    directions.remove(direction)
            if direction == "d":
                if row + i == max_size:
                    return True
                elif input[row][column] <= input[row + i][column]:
                    directions.remove(direction)
            if direction == "u":
                if row - i < 0:
                    return True
                elif input[row][column] <= input[row - i][column]:
                    directions.remove(direction)
        if not len(directions):
            return False
    # print("[" + str(row) + "," + str(column) + "] -> " + str(res))
    return False

def heigh_value(row, column, input, max_size):
    directions = ["r", "l", "u", "d"]
    res = 1
    for i in range(1, max_size):
        tmp_directions = directions.copy()
        for direction in tmp_directions:
            if direction == "r":
                if column + i == max_size:
                    res *= i - 1
                    directions.remove(direction)
                elif input[row][column] <= input[row][column + i]:
                    res *= i
                    directions.remove(direction)
            if direction == "l":
                if column - i < 0:
                    res *= i - 1
                    directions.remove(direction)
                elif input[row][column] <= input[row][column - i]:
                    res *= i
                    directions.remove(direction)
            if direction == "d":
                if row + i == max_size:
                    res *= i - 1
                    directions.remove(direction)
                elif input[row][column] <= input[row + i][column]:
                    res *= i
                    directions.remove(direction)
            if direction == "u":
                if row - i < 0:
                    res *= i - 1
                    directions.remove(direction)
                elif input[row][column] <= input[row - i][column]:
                    res *= i
                    directions.remove(direction)
        if not len(directions):
            break
    # print("[" + str(row) + "," + str(column) + "] -> " + str(res))
    return res


def read_file(file_name):
    res = []
    with open(file_name, "r") as f:
        for line in f.readlines():
            res.append([int(x) for x in line.strip()])
    return res


if __name__ == "__main__":
    input = read_file(FILE_NAME)
    print(process_input1(input))