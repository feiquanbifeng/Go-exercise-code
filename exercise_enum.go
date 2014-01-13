package main

import "fmt"

const (
  PRICE = iota
  SECOND = 1 << iota
)

func main() {
  fmt.Println(PRICE)
  fmt.Println(SECOND)
}
