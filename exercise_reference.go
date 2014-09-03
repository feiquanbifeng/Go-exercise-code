package main

import "fmt"
import "reflect"

func main() {
	p := [...]int{2, 3, 5, 7, 11, 13}
	s1 := p[1:3]
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(s1))
	
	changeArrayValue(p)
	fmt.Println(p)
	
	changeSliceValue(s1)
	fmt.Println(s1)
	
	fmt.Println(p)
}

func changeArrayValue(arr [6]int) {
	arr[0] = 100
}

func changeSliceValue(arr []int) {
	arr[0] = 100
}
