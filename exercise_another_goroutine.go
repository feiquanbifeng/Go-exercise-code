package main

import (
  "fmt"
  "runtime"
)

func filter(in, out chan int, prime int) {
  for {
    i := <- in
    if i % prime != 0 {
      out <- i
    }
  }
}

func main() {
  runtime.GOMAXPROCS(1)
  ch := make(chan int)

  go func(ch chan int) {
    for i := 2; ; i++ {
      ch <- i
    }
  }(ch)

  for {
    prime := <- ch
    fmt.Println(prime)

    ch1 := make(chan int)
    go filter(ch, ch1, prime)
    ch = ch1
  }
}
