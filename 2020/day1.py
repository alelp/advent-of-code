#!/usr/bin/env python

with open('input-day1.txt') as f:
    numbers = list(map(int, f.read().splitlines()))

    done1 = False
    done2 = False

    for x in numbers:
        for y in numbers:
            if not done1 and x + y == 2020:
                print("part 1", x * y)
                done1 = True

            if not done2:
                for z in numbers:
                    if x + y + z == 2020:
                        print("part 2", x * y * z)
                        done2 = True
                        break

            if done1 and done2:
                exit(0)

