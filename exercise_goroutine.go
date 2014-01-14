package main

import (
  "fmt"
  // "time"
  "runtime"
)

func main() {
  // c := make(chan bool)
  /*go func() {
    fmt.Println("Go Go !!")
    c <- true
    close(c)
  }()*/
  // <-c
  // time.Sleep(2 * time.Second)

  /*for v := range c {
    fmt.Println(v)
  }*/

  runtime.GOMAXPROCS(runtime.NumCPU())
  channel := make(chan bool, 10)
  
  for i := 0; i < 10; i++ {
    go Routine(channel)
  }
  
  for i := 0; i < 10; i++ {
    <- channel
  }
  
}

func Go() {
  fmt.Println("Go Go Go!!!")
}

func Routine(c chan bool) {
  a := 1
  for i := 0; i < 10000000; i++ {
    a += i
  }

  fmt.Println(a)
  
  // if index == 9 {
    c <- true
  // }
}
