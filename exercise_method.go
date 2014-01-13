package main

import "fmt"

type A struct {
  Name string
  age int
}

type B struct {
  Name string
}

type C int

func main() {
  a := A{}
  a.Print()
  fmt.Println(a.Name)
  // output the private val
  fmt.Println(a.age)

  b := B{}
  // method value
  b.Print()
  fmt.Println(b.Name)

  var c C
  // method expression
  (*C).Print(&c)
}

func (a *A) Print() {
  a.Name = "AA"
  a.age = 26
  fmt.Println("A")
}

func (b B) Print() {
  b.Name = "BB"
  fmt.Println("B")
}

func (c *C) Print() {
  fmt.Println("C")
}
