FILE_NAME = "input.txt"
START_SIZE = 14

def process_input(input):
    queue = []
    for i, c in enumerate(input):
        for iq, cq in enumerate(queue):
            if cq == c:
                queue = queue[iq+1:]
        queue.append(c)
        if len(queue) == START_SIZE:
            return i + 1
    return -1

def read_file(file_name):
    with open(file_name, "r") as f:
        return f.readlines()[0].strip()

if __name__ == "__main__":
    input = read_file(FILE_NAME)
    print(input)
    index = process_input(input)
    print(index)