package main

import (
  "fmt"
)

func main() {
  sm := make([]map[int] string, 5)
  for i := range sm {
    sm[i] = make(map[int] string)
    sm[i][1] = "OK"
    fmt.Println(sm[i])
  }

  fmt.Println(sm)
}
