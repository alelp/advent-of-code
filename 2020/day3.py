#!/usr/bin/env python

def get_tree_count(rows, slope_x, slope_y):
    tree_count = 0
    current_x = 0
    current_y = 0

    while current_y < len(rows):
        columns = rows[current_y]
        if columns[current_x] == '#':
            tree_count += 1

        current_x += slope_x
        if current_x >= len(columns):
            current_x -= len(columns)
        current_y += slope_y
    return tree_count


with open('input-day3.txt') as f:
    rows = f.read().splitlines()

    print("part 1:", get_tree_count(rows, 3, 1))

    slopes_list = [
        [1, 1],
        [3, 1],
        [5, 1],
        [7, 1],
        [1, 2],
    ]
    magick_number = 1

    for slope in slopes_list:
        magick_number *= get_tree_count(rows, slope[0], slope[1])

    print("part 2:", magick_number)
