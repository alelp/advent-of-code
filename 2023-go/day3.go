package main

import (
	"fmt"
	"strconv"
)

func main() {
  day := 3
  test_result1 := 4361
  test_result2 := 467835

  fmt.Println("---------------------------------------------------")
  fmt.Println("Day", day)
  fmt.Println("---------------------------------------------------")
  RunTest(day, 1, part1, test_result1);
  RunTest(day, 2, part2, test_result2);
  fmt.Println("---------------------------------------------------")
  RunDay(day, part1, part2);
  fmt.Println("---------------------------------------------------")
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

type Position struct {
  x int
  y int
}

func is_digit(c byte) bool {
  return '0' <= c && c <= '9'
}

func find_value(lines []string, x int, y int) (Position, int, bool) {
  var pos Position
  var value = -1
  var success = false

  if y >= 0 && y < len(lines) {
    line := lines[y]
    line_limit := len(line)

    if x >= 0 && x < len(line) {
      char := line[x]
      if is_digit(char) {
        pos.y = y
        pos.x = x

        for is_digit(line[pos.x]) && pos.x > 0 {
          pos.x--
        }

        if !is_digit(line[pos.x]) {
          pos.x++
        }

        var end = x

        for end < line_limit && is_digit(line[end]) {
          end++
        }

        value, _ = strconv.Atoi(line[pos.x:end])

        success = true
      }
    }
  }

  return pos, value, success
}

func part2(lines []string) int {
  var result = 0

  offsets := []int { -1, 0, 1 }

  for y, line := range lines {
    for x, char := range line {
      if char == '*' {
        var values = make(map[Position]int)

        for _, offset_x := range offsets {
          for _, offset_y := range offsets {
            position, value, success := find_value(lines, x + offset_x, y + offset_y)

            if success {
              values[position] = value
            }
          }
        }

        if len(values) > 2 {
          fmt.Println("ERROR: Too many numbers for gear at", x, y)
        } else if len(values) == 2 {
          var ratio = 1

          for _, value := range values {
            ratio *= value
          }

          result += ratio
        }
      }
    }
  }

  return result
}

