package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
  day := 2
  test_result1 := 8
  test_result2 := 2286

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
  limits := map[string]int {
    "red": 12,
    "green": 13,
    "blue": 14,
  }

  var result = 0

  for _, line := range lines {
    var temp = strings.Split(line, ":")
    if len(temp) < 2 { continue }
    game_id, _ := strconv.Atoi(strings.Split(temp[0], " ")[1])
    game_results := strings.Split(temp[1], ";")

    var valid_game = true

    for _, result := range game_results {
      draws := strings.Split(result, ", ")
      for _, draw := range draws {
        var color string
        var count int
        fmt.Sscanf(draw, "%d %s", &count, &color)

        if limits[color] < count {
          valid_game = false
        }
      }
    }

    if valid_game {
      result += game_id
    }
  }

  return result
}

func part2(lines []string) int {
  var result = 0

  for _, line := range lines {
    var temp = strings.Split(line, ":")
    if len(temp) < 2 { continue }
    game_results := strings.Split(temp[1], ";")
    limits := map[string]int {
      "red": 0,
      "green": 0,
      "blue": 0,
    }

    for _, result := range game_results {
      draws := strings.Split(result, ", ")
      for _, draw := range draws {
        var color string
        var count int
        fmt.Sscanf(draw, "%d %s", &count, &color)

        if limits[color] < count {
          limits[color] = count
        }
      }
    }

    var game_power = 1

    for _, count := range limits {
      game_power *= count
    }

    result += game_power
  }

  return result
}
