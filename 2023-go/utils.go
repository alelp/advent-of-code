
package main

import (
  "bufio"
  "fmt"
  "os"
 _ "strconv"
)

type solver_fn func([]string) int

func LoadInput(filename string) []string {
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

func RunTest(day int, part int, solver solver_fn, result int) {
  filename := fmt.Sprintf("tests/day%dpart%d.txt", day, part)
  test_lines := LoadInput(filename)
  fmt.Printf("Day %d part %d test passed: %t\n", day, part, solver(test_lines) == result)
}

func RunDay(day int, solver1 solver_fn, solver2 solver_fn) {
  filename := fmt.Sprintf("input/day%d.txt", day)
  var lines = LoadInput(filename);

  if lines == nil {
    fmt.Println("Error: Missing Day", day, "input.")
    return
  }

  fmt.Printf("Day %d part 1 result: %d\n", day, solver1(lines))
  fmt.Printf("Day %d part 2 result: %d\n", day, solver2(lines))
}
