#!/usr/bin/env python

def can_hold(desired_color, bag_color, bag_list):
    contents = bag_list[bag_color]
    if desired_color in contents:
        return True
    else:
        for content_colors in list(contents.keys()):
            if can_hold(desired_color, content_colors, bag_list):
                return True
        return False


def count_content(bag_color, bag_list):
    contents = bag_list[bag_color]
    count = 0
    for current_color, current_color_count in list(contents.items()):
        count += current_color_count
        count += current_color_count * count_content(current_color, bag_list)
    return count


with open('input-day7.txt') as f:
    rows = f.read().splitlines()

    bags = {}
    for row in rows:
        bag_color, contents = row.split(' bags contain ')
        raw_rules = contents[:-1].split(',')
        content_rules = {}
        if raw_rules[0].strip() != 'no other bags':
            for rule in raw_rules:
                count, color = rule.strip().split(' ', 1)
                count = int(count)
                if count == 1:
                    color = color.replace('bag', '').strip()
                else:
                    color = color.replace('bags', '').strip()

                content_rules[color] = count
        bags[bag_color] = content_rules

    target_color = 'shiny gold'
    count = 0
    for color in list(bags.keys()):
        if can_hold(target_color, color, bags):
            count += 1

    print("part 1:", count)
    print("part 2:", count_content(target_color, bags))

