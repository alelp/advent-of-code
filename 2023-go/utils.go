
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

func LoadInput(day int) []string {
  filename := fmt.Sprintf("input/day%d.txt", day)

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

func RunTests(day int, part1 DayPart, part2 DayPart) {
  fmt.Println("-----------------")
  fmt.Printf("Day %d part 1 test passed: %t\n", day, part1.solver(part1.test_lines) == part1.test_result)
  fmt.Printf("Day %d part 2 test passed: %t\n", day, part2.solver(part2.test_lines) == part2.test_result)
  fmt.Println("-----------------")
}

func RunDay(day int, part1 DayPart, part2 DayPart) {
  var lines = LoadInput(day);

  if lines == nil {
    fmt.Println("Error: Missing Day", day, "input.")
    return
  }

  fmt.Println("-----------------")
  fmt.Printf("Day %d part 1 result: %d\n", day, part1.solver(lines))
  fmt.Printf("Day %d part 2 result: %d\n", day, part2.solver(lines))
  fmt.Println("-----------------")
}
