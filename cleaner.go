package main

import (
  "os"
  "bufio"
  "fmt"
)

func parse(file *os.File) error {
  lines := map[string][]string{}
  scanner := bufio.NewScanner(file)
  selector := ""
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      continue
    }
    if line[0] == '#' || line[0] == '.' {
      selector = line[0:len(line) - 2]
      if _, ok := lines[selector]; !ok {
        lines[selector] = []string{}
      }
    } else if line[0] != '}' {
      lines[selector] = append(lines[selector], line)
    }
  }

  fmt.Println(lines)

  return scanner.Err()
}

func main() {
  file, e := os.Open("test.scss");
  if e != nil {
    panic(e)
  }
  defer file.Close()

  if err := parse(file); err != nil {
    panic(err)
  }
}
