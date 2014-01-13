package main

import (
  "fmt"
  "reflect"
)

type User struct {
  Id int
  Name string
  Age int
}

type Manager struct {
  User
  title string
}

func (u User) Hello() {
  fmt.Println("Hello world.")
}

func (u User) Say(name string, num int) (result string) {
  // fmt.Println("My name:", name, " Address:", address, " prive name", u.Name)
  fmt.Println("The call name:", name, " User Name:", u.Name, " NUM=", num)
  return "Successful!"
}

func main() {
  u := User{1, "OK", 12}
  Info(u)

  m := Manager{User: User{1, "HaHa", 22}, title: "132"}
  t := reflect.TypeOf(m)

  fmt.Println("%#v\n", t.Field(1))

  x := 123
  v := reflect.ValueOf(&x)
  v.Elem().SetInt(999)
 
  fmt.Println(x)

  vu := reflect.ValueOf(u)
  vm := vu.MethodByName("Say")

  args := []reflect.Value{reflect.ValueOf("Hello"), reflect.ValueOf(55)}
  ret := vm.Call(args) 
  fmt.Println(ret) 
}

func Info(o interface{}) {
  t := reflect.TypeOf(o)
  
  fmt.Println("Type:", t.Name())

  v := reflect.ValueOf(o)
  fmt.Println("Fileds:")

  for i := 0; i < t.NumField(); i++ {
    f := t.Field(i)
    val := v.Field(i).Interface()
    fmt.Println("%6s: %v = %v\n", f.Name, f.Type, val)
  }
 
  for i := 0; i < t.NumMethod(); i++ {
    m := t.Method(i)
    fmt.Println("%6s: %v\n", m.Name, m.Type)
  }
}
