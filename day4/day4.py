def solution_first_part(path: str) -> None:
    score = 0
    with open(file=path, newline="\n") as input_data:
        for row in input_data:
            row = row.rstrip("\n")
            groups = row.split(",")
            groups = [x.split("-") for x in groups]
            if (
                int(groups[0][0]) >= int(groups[1][0])
                and int(groups[0][1]) <= int(groups[1][1])
            ) or (
                int(groups[0][0]) <= int(groups[1][0])
                and int(groups[0][1]) >= int(groups[1][1])
            ):
                score += 1
            else:
                continue

    print(score)


def solution_second_part(path: str) -> None:
    score = 0
    with open(file=path, newline="\n") as input_data:
        for row in input_data:
            row = row.rstrip("\n")
            groups = row.split(",")
            groups = [x.split("-") for x in groups]
            if (
                (
                    int(groups[0][0]) >= int(groups[1][0])
                    and int(groups[0][1]) <= int(groups[1][1])
                )
                or (
                    int(groups[0][0]) <= int(groups[1][0])
                    and int(groups[0][1]) >= int(groups[1][1])
                )
                or int(groups[0][1]) == int(groups[1][0])
                or (
                    int(groups[0][1]) > int(groups[1][0])
                    and int(groups[0][1]) < int(groups[1][1])
                )
                or (
                    int(groups[0][0]) < int(groups[1][1])
                    and int(groups[0][1]) > int(groups[1][1])
                )
                or int(groups[0][0]) == int(groups[1][1])
            ):
                score += 1
            else:
                continue

    print(score)


solution_first_part("input4.txt")
solution_second_part("input4.txt")
