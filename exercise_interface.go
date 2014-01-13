package main

import "fmt"


type USB interface {
  Name() string
  Connecter
}

type Connecter interface {
  Connect()
}

type PhoneConnect struct {
  name string
}

func (pc PhoneConnect) Name() string {
  return pc.name
}

func (pc PhoneConnect) Connect() {
  fmt.Println("Connect:", pc.name)
}

func main() {
  var a USB
  a = PhoneConnect{"PhoneConnect"}

  a.Connect()
  Disconnect(a)
}

func Disconnect(usb USB) {
  if pc, ok := usb.(PhoneConnect); ok {
    fmt.Println("Disconnected:", pc.name)
    return
  }
  fmt.Println("Unknown device")
}
