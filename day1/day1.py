def solution_first_part(input):
    hash_map = {}
    counter = 0
    with open(file=input, newline="") as input_data:
        for row in input_data:
            if row.rstrip("\n").isnumeric():
                if counter in hash_map:
                    hash_map[counter] += int(row)
                else:
                    hash_map[counter] = 0
                    hash_map[counter] += int(row)
            elif row.rstrip("\n") == "":
                counter += 1
            else:
                break

    solution = float("-inf")

    for _, value in hash_map.items():
        if value > solution:
            solution = value

    print(solution)


def solution_second_part(input):
    hash_map = {}
    counter = 0
    with open(file=input, newline="") as input_data:
        for row in input_data:
            if row.rstrip("\n").isnumeric():
                if counter in hash_map:
                    hash_map[counter] += int(row)
                else:
                    hash_map[counter] = 0
                    hash_map[counter] += int(row)
            elif row.rstrip("\n") == "":
                counter += 1
            else:
                break

    largest_three = [float("-inf")] * 3

    for _, value in hash_map.items():
        if value > largest_three[0]:
            largest_three[0] = value
            largest_three.sort()

    print(sum(largest_three))


print(solution_first_part("input1.txt"))
print(solution_second_part("input1.txt"))
