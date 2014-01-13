package main

import (
  "fmt"
  "sort"
)

func main() {
  m := map[int] string {1: "a", 6: "b", 2: "c", 4: "d", 5: "e"}
  s := make([]int, len(m))
  i := 0

  for k, _ := range m {
    s[i] = k
    i++
  }

  sort.Ints(s)
  fmt.Println(s)
}
