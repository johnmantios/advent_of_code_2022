from itertools import islice

priorities = {
    "a": 1,
    "b": 2,
    "c": 3,
    "d": 4,
    "e": 5,
    "f": 6,
    "g": 7,
    "h": 8,
    "i": 9,
    "j": 10,
    "k": 11,
    "l": 12,
    "m": 13,
    "n": 14,
    "o": 15,
    "p": 16,
    "q": 17,
    "r": 18,
    "s": 19,
    "t": 20,
    "u": 21,
    "v": 22,
    "w": 23,
    "x": 24,
    "y": 25,
    "z": 26,
    "A": 27,
    "B": 28,
    "C": 29,
    "D": 30,
    "E": 31,
    "F": 32,
    "G": 33,
    "H": 34,
    "I": 35,
    "J": 36,
    "K": 37,
    "L": 38,
    "M": 39,
    "N": 40,
    "O": 41,
    "P": 42,
    "Q": 43,
    "R": 44,
    "S": 45,
    "T": 46,
    "U": 47,
    "V": 48,
    "W": 49,
    "X": 50,
    "Y": 51,
    "Z": 52,
}


def solution_first_part(input: str) -> None:

    score = 0

    with open(file=input, newline="\n") as input_data:
        for row in input_data:
            row = row.rstrip("\n")
            first_half_set = set(row[: len(row) // 2])
            second_half_set = set(row[len(row) // 2 :])
            common_item = first_half_set.intersection(second_half_set)
            score += priorities[min(common_item)]

    print(score)


def solution_second_part(input: str) -> None:
    score = 0

    with open(file=input, newline="\n") as input_data:
        while True:
            group = list(islice(input_data, 3))  # read groups of 3 lines
            group = [x.rstrip("\n") for x in group]  # strip the \n
            group = [set(x) for x in group]  # turn each element in the list to a set
            if group:
                score += priorities[min(set.intersection(*group))]
            else:
                break

    print(score)


solution_first_part("input3.txt")
solution_second_part("input3.txt")
