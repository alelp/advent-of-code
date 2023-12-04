package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
  part1 := DayPart{
    []string {
      "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
      "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
      "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
      "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    },
    8,
    part1,
  }

  part2 := DayPart{
    []string {
      "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
      "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
      "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
      "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
      "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    },
    2286,
    part2,
  }

  day := 2
  RunTests(day, part1, part2);
  RunDay(day, part1, part2);
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
