package main

import (
    "fmt"
)

type SellInterface interface {
    Sell() interface{}
}

type RedWineFactory struct {

}

type RedWineProxy struct {
    factory RedWineFactory
    sellCount int
}

func (self *RedWineFactory) Sell() interface{} {
    fmt.Println("This is real role")
    return nil
}

func (self *RedWineProxy) Sell() interface{} {
    obj := self.factory.Sell()
    self.sellCount++
    fmt.Println("The factory selled :", self.sellCount)
    return obj
}

func main() {
    var sell SellInterface = &RedWineProxy {sellCount: 4}
    sell.Sell()
}
