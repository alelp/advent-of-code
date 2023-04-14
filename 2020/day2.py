#!/usr/bin/env python

with open('input-day2.txt') as f:
    passwords = f.read().splitlines()

    valid = 0
    for password_data in passwords:
        limits, letter, password = [x.strip(' :') for x in password_data.split(' ')]
        limit_min, limit_max = [int(x) for x in limits.split('-')]
        count = password.count(letter);
        if count >= limit_min and count <= limit_max:
            valid += 1
    print('part 1:', valid)

    valid = 0
    for password_data in passwords:
        limits, letter, password = [x.strip(' :') for x in password_data.split(' ')]
        pos_1, pos_2 = [int(x) - 1 for x in limits.split('-')]
        if password[pos_1] == letter and password[pos_2] != letter:
            valid += 1
        elif password[pos_1] != letter and password[pos_2] == letter:
            valid += 1
    print('part 2:', valid)


