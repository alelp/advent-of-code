package main

import (
	"fmt"
	"strconv"
	_ "strings"
)

func main() {
  part1 := DayPart{
    []string {
      "467....114",
      "...*......",
      "..35...633",
      "......#...",
      "617*......",
      ".....+.58.",
      "..592.....",
      "......755.",
      "...$.*....",
      ".664.598..",
    },
    4361,
    part1,
  }

  part2 := DayPart{
    []string {
    },
    -1,
    part2,
  }

  day := 3
  fmt.Println("Day", day)
  RunTests(day, part1, part2);
  RunDay(day, part1, part2);
}

func part1(lines []string) int {
  var result = 0

  for y, line := range lines {
    line_limit := len(line) - 1
    var begin = -1
    for x, char := range line {
      is_digit := '0' <= char && char <= '9'
      line_ended := x == line_limit

      if is_digit && begin == -1 {
          begin = x
      } 

      if begin != -1 && (!is_digit || line_ended) {
        end := x
        if line_ended && is_digit {
          end++
        }

        num, _ := strconv.Atoi(line[begin:end])
        //fmt.Println("Found num:", num)

        if is_symbol(lines, begin-1, y) || is_symbol(lines, x, y) {
          goto ADD_NUM
        }

        for i := begin -1; i <= x; i++ {
          if is_symbol(lines, i, y-1) || is_symbol(lines, i, y+1) {
            goto ADD_NUM
          }
        }

        goto END
ADD_NUM:
        result += num
END:
        begin = -1
      }
    }
  }

  return result
}

func is_symbol(lines []string, x int, y int) bool {
  if y < 0 || y >= len(lines) {
    return false
  }

  line := lines[y]

  if x < 0 || x >= len(line) {
    return false
  }

  char := line[x]
  /*
  if char != '.' {
    fmt.Printf("%c - %t\n", char, char < '0' || char > '9')
  }
  */
  return char != '.' && (char < '0' || char > '9')
}

func part2(lines []string) int {
  var result = 0
  return result
}
