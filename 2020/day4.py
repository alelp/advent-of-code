#!/usr/bin/env python

def validate_height(x):
    unit = x[-2:]
    try:
        number = int(x[0:-2])
    except Exception:
        number = 0

    if unit == 'cm':
        return number >= 150 and number <= 193
    elif unit == 'in':
        return number >= 59 and number <= 76
    else:
        return False


req_fields = {
    # (Birth Year)
    "byr": lambda x: int(x) >= 1920 and int(x) <= 2002,

    # (Issue Year)
    "iyr": lambda x: int(x) >= 2010 and int(x) <= 2020,

    # (Expiration Year)
    "eyr": lambda x: int(x) >= 2020 and int(x) <= 2030,

    # (Height)
    "hgt": validate_height,

    # (Hair Color)
    "hcl": lambda x: x[0] == '#' and int(x[1:-1], 16),

    # (Eye Color)
    "ecl": lambda x: x in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"],

    # (Passport ID)
    "pid": lambda x: len(x) == 9,
}

ignore_fields = {
    # (Country ID)
    "cid": lambda x: True,
}


def count_valids(passports, field_validation):
    valid_passports = 0
    for current_passport in passports:
        is_valid = True
        if all(field in current_passport for field in list(req_fields.keys())):
            if field_validation:
                for key, value in list(current_passport.items()):
                    if key in ignore_fields:
                        continue
                    elif not req_fields[key](value):
                        is_valid = False
        else:
            is_valid = False

        if is_valid:
            valid_passports += 1

    return valid_passports


with open('input-day4.txt') as f:
    rows = f.read().splitlines()
    rows.append("")  # make sure the last passport is being processed

    passports = []
    current_passport = {}
    for row in rows:
        if len(row) == 0:
            passports.append(current_passport)
            current_passport = {}
        else:
            data = row.split(' ')
            for field in data:
                key, value = field.split(':')
                current_passport[key] = value

    print("part 1:", count_valids(passports, False))
    print("part 2:", count_valids(passports, True))

