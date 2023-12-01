#!/usr/bin/env python

def snafu_to_decimal(snafu):
    decimal = 0
    base = 1
    for i in range(len(snafu)-1, -1, -1):
        if snafu[i] == '=':
            decimal += (-2) * base
        elif snafu[i] == '-':
            decimal += (-1) * base
        elif snafu[i] == '0':
            decimal += 0 * base
        elif snafu[i] == '1':
            decimal += 1 * base
        elif snafu[i] == '2':
            decimal += 2 * base
        base *= 5
    return decimal

snafu_list = []
with open('input-25.txt', 'r') as file:
    for line in file:
        snafu_list.append(line.strip())

decimal_sum = 0
for snafu in snafu_list:
    decimal_sum += snafu_to_decimal(snafu)

print(decimal_sum)

