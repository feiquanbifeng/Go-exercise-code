package main

import "fmt"

func main() {
  LABEL1:
  for i := 0; i < 10; i++ {
      for {
        fmt.Println(i)
        continue LABEL1
      }
   }
   fmt.Println("OK")
}
