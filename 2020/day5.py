#!/usr/bin/env python

import math
import sys

debug = 0
if debug:
    seats_data = [
        "BFFFBBFRRR",  #: row 70, column 7, seat ID 567.
        "FFFBBBFRRR",  #: row 14, column 7, seat ID 119.
        "BBFFBBFRLL",  #: row 102, column 4, seat ID 820.
    ]
else:
    with open('input-day5.txt') as f:
        seats_data = f.read().splitlines()

highest_id = 0
lowest_id = sys.maxsize
all_ids = []
for seat in seats_data:
    row_min = 0
    row_max = 127

    col_min = 0
    col_max = 7

    for c in seat:
        if c == 'F':
            row_max -= math.ceil((row_max - row_min) / 2.0)
        elif c == 'B':
            row_min += math.ceil((row_max - row_min) / 2.0)
        elif c == 'L':
            col_max -= math.ceil((col_max - col_min) / 2.0)
        elif c == 'R':
            col_min += math.ceil((col_max - col_min) / 2.0)
        # print(c, row_min, row_max, col_min, col_max)

    seat_id = row_max * 8 + col_max
    # print( seat, row_max, col_max, seat_id)
    # print( '------')
    all_ids.append(seat_id)
    if seat_id > highest_id:
        highest_id = int(seat_id)
    if seat_id < lowest_id:
        lowest_id = int(seat_id)

print("part 1:\n\t", highest_id)

print("\npart 2:")
for i in range(lowest_id, highest_id):
    if i not in all_ids and (i + 1) in all_ids and (i - 1) in all_ids:
        print("\t id candidate:", i)

