package main

import "fmt"

func main() {
    const (
        A1 = iota
        A2
        str = "hello"
        s
        A3 = iota
        A4
    )
    /* A1= 0 A2= 1 s= hello str= hello A3= 4 A4= 5 */
    fmt.Println("A1=", A1, "A2=", A2, "s=", s, "str=", str, "A3=", A3, "A4=", A4)
}
