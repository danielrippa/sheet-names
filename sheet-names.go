package main

import (
  "os"
  "fmt"
  "github.com/xuri/excelize/v2"
)

func main() {
  args := os.Args[1:]

  f, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, name := range f.GetSheetList() {
    fmt.Println(name)
  }
}
