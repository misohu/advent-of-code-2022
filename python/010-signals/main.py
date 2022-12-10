from dataclasses import dataclass

FILE_NAME = "input.txt"

@dataclass
class System:
    register: int
    cycles: int
    result: int
    targets: list[int]


def process_input(file_name):
    s = System(1, 0, 0, range(19, 225, 40))
    with open(file_name) as f:
        for line in f:
            match line.strip().split():
                case "addx", value:
                    s = process_register(s, 2, int(value))
                case _:
                    s = process_register(s, 1, 0)   
    return s

def process_register(s: System, c: int, val: int):
    for i in range(c):
        if s.cycles in s.targets:
            s.result += (s.register * (s.cycles+1))
        if i == c-1:
            s.register += val
        s.cycles+=1
    return s

if __name__ == "__main__":
    s = process_input(FILE_NAME)
    print(s.result)
    


        