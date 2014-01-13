package main

import (
  "fmt"
)

func main() {
  // var m map[int] string
  // m = map[int] string {}

  // var m = make(map[int] string)
  // m := make(map[int] string)
  m := make(map[int] map[int] string)
  // m[1] = make(map[int] string)
  a, ok := m[2][1]

  if !ok {
    m[2] = make(map[int] string)
  }
  m[2][1] = "GOOD"
  // delete(m, 1)
  // a := m[1][1]
  a, ok = m[2][1]
  fmt.Println(a, ok)
}
