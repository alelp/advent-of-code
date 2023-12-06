package main

import (
	"fmt"
	"strings"
)

func part1(lines []string) int {
  var result = 0

  for _, line := range lines {
    card_parts := strings.Split(line, ":")
    numbers := strings.Split(card_parts[1], "|")
    winning_numbers := strings.Split(numbers[0], " ")
    card_numbers := strings.Split(numbers[1], " ")

    var card_score = 0

    for _, num := range card_numbers {
      for _, win_num := range winning_numbers {
        if num == win_num && num != "" {
          if card_score == 0 {
            card_score = 1
          } else {
            card_score *= 2
          }

          break;
        }
      }
    }

    result += card_score
  }

  return result
}

func part2(lines []string) int {
  var result = -1
  return result
}

func main() {
  day := 4
  test_result1 := 13
  test_result2 := 30

  fmt.Println("---------------------------------------------------")
  fmt.Println("Day", day)
  fmt.Println("---------------------------------------------------")
  RunTest(day, 1, part1, test_result1);
  RunTest(day, 2, part2, test_result2);
  fmt.Println("---------------------------------------------------")
  RunDay(day, part1, part2);
  fmt.Println("---------------------------------------------------")
}
