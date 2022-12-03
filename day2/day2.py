import csv


def solution_first_part(input):

    me = {"X": 1, "Y": 2, "Z": 3}

    score = 0

    with open(file=input, newline="") as input_data:
        for row in csv.reader(input_data, delimiter=" "):
            if row[1] == "X" and row[0] == "A":
                score += me[row[1]] + 3
            elif row[1] == "X" and row[0] == "B":
                score += me[row[1]]
            elif row[1] == "X" and row[0] == "C":
                score += me[row[1]] + 6
            elif row[1] == "Y" and row[0] == "A":
                score += me[row[1]] + 6
            elif row[1] == "Y" and row[0] == "B":
                score += me[row[1]] + 3
            elif row[1] == "Y" and row[0] == "C":
                score += me[row[1]]
            elif row[1] == "Z" and row[0] == "A":
                score += me[row[1]]
            elif row[1] == "Z" and row[0] == "B":
                score += me[row[1]] + 6
            elif row[1] == "Z" and row[0] == "C":
                score += me[row[1]] + 3

    print(score)


def solution_second_part(input):

    points = {"X": 1, "Y": 2, "Z": 3}
    lose = {"A": "Z", "B": "X", "C": "Y"}
    tie = {"A": "X", "B": "Y", "C": "Z"}
    win = {"A": "Y", "B": "Z", "C": "X"}

    score = 0

    with open(file=input, newline="") as input_data:
        for row in csv.reader(input_data, delimiter=" "):
            if row[1] == "X":
                score += points[lose[row[0]]]
            elif row[1] == "Y":
                score += points[tie[row[0]]] + 3
            elif row[1] == "Z":
                score += points[win[row[0]]] + 6

    print(score)


print(solution_first_part("input2.txt"))
print(solution_second_part("input2.txt"))
