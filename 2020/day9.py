#!/usr/bin/env python

def is_valid(number, preamble):
    if len(preamble) <= 25:
        return True
    else:
        for i in range(len(preamble)):
            for j in range(len(preamble)):
                if i == j:
                    continue
                x = preamble[i]
                y = preamble[j]
                if x + y == number:
                    return True
        return False


def find_first_invalid(numbers, debug_print=False):
    for i in range(len(numbers)):
        current_number = numbers[i]
        preamble = numbers[:i]
        if is_valid(current_number, preamble):
            if debug_print:
                print(i, current_number)
            continue
        else:
            if debug_print:
                print('invalid number', current_number, 'at index ', i)
            return current_number, i


with open('input-day9.txt') as f:
    numbers = [int(x) for x in f.read().splitlines()]

    first_invalid_number, first_invalid_index = find_first_invalid(numbers)

    print('part 1:', first_invalid_number)

    lower_bound = 0
    upper_bound = 0
    weak_range = [0]

    while upper_bound < first_invalid_index:
        range_sum = 0
        for i in range(lower_bound, upper_bound):
            range_sum += numbers[i]

        if range_sum > first_invalid_number:
            lower_bound += 1
        elif range_sum < first_invalid_number:
            upper_bound += 1
        else:
            weak_range = numbers[lower_bound:upper_bound]
            print('range found!', weak_range)
            break

    print('part 2:', min(weak_range) + max(weak_range))

