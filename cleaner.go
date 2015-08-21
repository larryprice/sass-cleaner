package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
)

type Style struct {
  Attribute string
  Value string
}

func NewStyle(fromLine string) Style {
  vals := strings.SplitN(fromLine, ":", 2)
  return Style{
    Attribute: strings.TrimSpace(vals[0]),
    Value: strings.TrimSpace(vals[1]),
  }
}

type Styles map[string]string

func Parse(file *os.File) error {
  lines := map[string]Styles{}
  scanner := bufio.NewScanner(file)
  var selector string

  for scanner.Scan() {
    line := strings.TrimRight(strings.TrimSpace(strings.Replace(scanner.Text(), "{", "", 1)), ";")
    if line == "" || line == "}" {
      continue
    }

    if line[0] == '#' || line[0] == '.' {
      selector = line
      if _, ok := lines[selector]; !ok {
        lines[selector] = Styles{}
      }
      continue
    }

    style := NewStyle(line)
    lines[selector][style.Attribute] = style.Value
  }

  fmt.Println(lines)

  return scanner.Err()
}

func main() {
  file, e := os.Open("test/test.scss");
  if e != nil {
    panic(e)
  }
  defer file.Close()

  if err := Parse(file); err != nil {
    panic(err)
  }
}
