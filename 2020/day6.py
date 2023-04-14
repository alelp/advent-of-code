#!/usr/bin/env python

with open('input-day6.txt') as f:
    rows = f.read().splitlines()
    rows.append("")  # make sure the last group is being processed

    total_answers = 0
    current_group = set()
    for row in rows:
        if len(row) == 0:
            total_answers += len(current_group)
            current_group = set()
        else:
            for answer in row:
                current_group.add(answer)

    print("part 1:", total_answers)

    total_answers = 0
    current_group = []
    for row in rows:
        if len(row) == 0:
            x = current_group[0].intersection(*current_group[1:])
            total_answers += len(x)
            current_group = []
        else:
            current_group.append(set(row))

    print("part 2:", total_answers)

