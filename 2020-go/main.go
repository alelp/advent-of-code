
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Provide some day numbers to process!")
    return;
  }

  functionMap := map[string]func([]string){
    "1": day1,
    "2": day2,
  }

  for _,value := range os.Args[1:] {
    filename := "input-day" + value + ".txt"
    fmt.Println("Process day", value)

    readFile, err := os.Open(filename)

    if err != nil {
      fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()

    functionMap[value](fileLines);
  }
}

func day1(fileLines []string) {
  for i, line1 := range fileLines {
    for j, line2 := range fileLines[i+1:] {
      int1, _ := strconv.Atoi(line1)
      int2, _ := strconv.Atoi(line2)
      if i != j && int1 + int2 == 2020 {
        fmt.Println("Day Part 1:", int1 * int2)
        goto day1_step2  
      }
    }
  }

day1_step2:
  for i, line1 := range fileLines {
    for j, line2 := range fileLines[i+1:] {
      for k, line3 := range fileLines[j+1:] {
        int1, _ := strconv.Atoi(line1)
        int2, _ := strconv.Atoi(line2)
        int3, _ := strconv.Atoi(line3)
        if i != j && j != k && int1 + int2 + int3 == 2020 {
          fmt.Println("Day Part 2:", int1 * int2 * int3)
          goto day1_end
        }
      }
    }
  }

day1_end:
}

func day2(fileLines []string) {
  for i, line := range fileLines {
    fmt.Println(i+1, line)
  }
}

