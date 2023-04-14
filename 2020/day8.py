#!/usr/bin/env python

def safe_run(file_name, print_lines=False):
    with open(file_name) as f:
        instructions = f.read().splitlines()

        acc = 0
        op_history = []
        op_count = len(instructions)
        i = 0
        while i < op_count:
            if i in op_history:
                print('ERROR! loop at line', i)
                return acc

            op_history.append(i)
            line = instructions[i]
            if print_lines:
                print(i, line)
            op, value = line.split(' ')

            if op == 'jmp':
                line_no = i
                if value[0] == '+':
                    i += int(value[1:])
                elif value[0] == '-':
                    i -= int(value[1:])
                if i < 0 or i >= op_count:
                    print('ERROR! out of bounds instruction', i, 'at line', line_no)
                    return acc
            else:
                i += 1
                if op == 'nop':
                    continue
                elif op == 'acc':
                    if value[0] == '+':
                        acc += int(value[1:])
                    elif value[0] == '-':
                        acc -= int(value[1:])
        return acc


print("part 1", safe_run('input-day8.txt'))
print("part 2", safe_run('input-day8-fix.txt', print_lines=True))

