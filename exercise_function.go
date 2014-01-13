package main

import (
  "fmt"
)

func main() {
  A(1, 2, 3)
  // the function is also value type
  b := B
  b()
  
  c := func() {
    fmt.Println("anonymity c")
  }
  c()

  // closure attention the scope
  f := closure(10)
  fmt.Println(f(1))
  fmt.Println(f(2))

  // the destruct function
  fmt.Println("a")
  defer fmt.Println("b")
  defer fmt.Println("c")
  
  for i := 0; i < 3; i++ {
    defer (func(i int) {
      fmt.Println(i)
    })(i)
  }

  C()
}

func A(a ...int) {
  fmt.Println(a)
}

func B() {
  fmt.Println("Function B")
}

func closure(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}

func C() {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("Recover in C")
    }
  }()
  panic("Panic C")
}
