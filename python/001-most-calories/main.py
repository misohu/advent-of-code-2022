FILE_NAME = "input.txt"

def load_data(file_name):
    file = open(file_name, "r")
    lines = file.readlines()
    return lines

def process_elfs(lines):
    elfs = []
    elf = []

    for line in lines:
        if line == '\n':
           elfs += [sum(elf)] 
           elf = []
           continue
        elf += [int(line)]
    elfs += [sum(elf)] 
    return elfs

if __name__ == "__main__":
    lines = load_data(FILE_NAME)
    elfs = process_elfs(lines)
    elfs = sorted(elfs, reverse=True)
    print(sum(elfs[:3]))