package main

import (
  "fmt"
)

func main() {
  for i := 0; i < 3; i++ {
    v := 1
    fmt.Println(&v)
  }
}
