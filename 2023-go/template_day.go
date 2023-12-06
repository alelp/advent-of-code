package main

import (
 "fmt"
)

func part1(lines []string) int {
  var result = -1
  return result
}

func part2(lines []string) int {
  var result = -1
  return result
}

func main() {
  day := 0
  test_result1 := -1
  test_result2 := -1

  fmt.Println("---------------------------------------------------")
  fmt.Println("Day", day)
  fmt.Println("---------------------------------------------------")
  RunTest(day, 1, part1, test_result1);
  RunTest(day, 2, part2, test_result2);
  fmt.Println("---------------------------------------------------")
  RunDay(day, part1, part2);
  fmt.Println("---------------------------------------------------")
}
