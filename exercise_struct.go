package main

import (
  "fmt"
)

type Person struct {
  Name string
  Age int
  string
  Contact struct {
    Phone int
    Address string
  }
}

func main() {
  // a := Person {"MY", 26}
  a := Person {
      Name: "MY",
      Age: 26,
  }
  // a.Name = "JY"
  // a.Age = 27
  fmt.Println(a)
  a.Contact.Phone = 2344
  A(&a)
  fmt.Println(a)
}

func A(p *Person) {
  p.Age = 13
  fmt.Println("A", p)
}
