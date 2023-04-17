#!/usr/bin/env python

with open('input-day2.txt') as f:
    passwords = f.read().splitlines()

    check1_valids = 0
    check2_valids = 0
    for password_data in passwords:
        limits, letter, password = [x.strip(' :') for x in password_data.split(' ')]
        limit_min, limit_max = [int(x) for x in limits.split('-')]
        count = password.count(letter)

        if count >= limit_min and count <= limit_max:
            check1_valids += 1

        pos_1, pos_2 = limit_min - 1, limit_max - 1
        if (password[pos_1] == letter and password[pos_2] != letter) or (password[pos_1] != letter and password[pos_2] == letter):
            check2_valids += 1

    print('part 1:', check1_valids)
    print('part 2:', check2_valids)

