FILE_NAME = "input.txt"

def read_file(file_name):
    result = []
    with open(file_name, "r") as f:
        for line in f:
            result.append(line.strip())
    return result

def process_bags(bags):
    result = 0
    for bag in bags:
        result += encode_char(process_bag(bag))
    return result

def process_bags2(bags, number_in_group = 3):
    i = 0
    result = 0
    while i < len(bags):
        e = process_group(bags[i:i+number_in_group])
        result += encode_char(e)
        i+=number_in_group
    return result

def process_group(group, number_in_group = 3):
    for e in group[0]:
        if e in group[1] and e in group[2]:
            return e

def process_bag(bag):
    middle = int(len(bag)/2)
    for e in bag[:middle]:
        if e in bag[middle:]:
            return e

def encode_char(charc):
    if ord(charc) >= ord('a'):
        return ord(charc) - ord('a') + 1
    else:
        return ord(charc) - ord('A') + 27

if __name__ == "__main__":
    bags = read_file(FILE_NAME)
    # print(process_bags(bags))
    print(process_bags2(bags))