package main

import (
  "fmt"
)

type Person struct {
  GetName func(n string) string
}

func (self Person) Say() {
  fmt.Println("say person")
}

func Hello(n string) string {
  return n
}

func main() {
  p := &Person {Hello}
  ret := p.GetName("good")
  fmt.Println(ret)
  p.Say()
}
