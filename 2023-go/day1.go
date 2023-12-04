package main

import (
	"strconv"
	"strings"
)

func main() {
  part1 := DayPart{
    []string {
      "1abc2", 
      "pqr3stu8vwx", 
      "a1b2c3d4e5f", 
      "treb7uchet", 
    },
    142,
    part1,
  }

  part2 := DayPart{
    []string {
      "two1nine",
      "eightwothree",
      "abcone2threexyz",
      "xtwone3four",
      "4nineeightseven2",
      "zoneight234",
      "7pqrstsixteen",
    },
    281,
    part2,
  }

  RunDay("1", part1, part2);
}

func part2(lines []string) int {
  digits := map[string]int {
    "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
    "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
    "seven": 7, "eight": 8, "nine": 9,
  }

  var total_sum = 0

  for _, line := range lines {
    var idx1 = len(line)
    var val1 = -1
    var idx2 = -1
    var val2 = -1

    for key, val := range digits {
      check1 := strings.Index(line, key)
      if check1 >= 0 && check1 <= idx1 {
        idx1 = check1
        val1 = val
      }

      check2 := strings.LastIndex(line, key)
      if check2 >= 0 && check2 > idx2 {
        idx2 = check2
        val2 = val
      }
    }

    total_sum += val1 * 10 + val2
  }

  return total_sum
}

func part1(lines []string) int {
  var total_sum = 0

  for _, line := range lines {
    var char1 = 'A'
    var char2 = 'B'

    for _, char := range line {
      if '0' <= char && char <= '9' {
        if char1 == 'A' {
          char1 = char;
          char2 = char;
        } else {
          char2 = char;
        }
      }
    }

    line_val, _ := strconv.Atoi(string(char1) + string(char2))
    total_sum += line_val
  }

  return total_sum
}
