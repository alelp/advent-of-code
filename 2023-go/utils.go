
package main

import (
  "bufio"
  "fmt"
  "os"
 _ "strconv"
)

type DayPart struct {
  test_lines []string
  test_result int
  solver func([]string) int
}

func LoadInput(day string) []string {
  filename := "input/day" + day + ".txt"

  input_file, err := os.Open(filename)

  if err != nil {
    fmt.Println("No input file,", filename)
    return nil
  }

  fileScanner := bufio.NewScanner(input_file)

  fileScanner.Split(bufio.ScanLines)

  var fileLines []string

  for fileScanner.Scan() {
    fileLines = append(fileLines, fileScanner.Text())
  }

  input_file.Close()

  return fileLines;
}

func RunDay(day string, part1 DayPart, part2 DayPart) {
  var lines = LoadInput(day);

  if lines == nil {
    fmt.Println("Error: Missing Day", day, "input.")
    return
  }

  fmt.Println("-----------------")
  fmt.Println("Day", day, "part 1:")
  fmt.Printf("Test passed: %t\n", part1.solver(part1.test_lines) == part1.test_result)
  fmt.Printf("Result: %d\n", part1.solver(lines))
  fmt.Println("-----------------")
  fmt.Println("Day", day, "part 2:")
  fmt.Printf("Test passed: %t\n", part2.solver(part2.test_lines) == part2.test_result)
  fmt.Printf("Result: %d\n", part2.solver(lines))
  fmt.Println("-----------------")
}
