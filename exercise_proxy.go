package main

import (
    "fmt"
)

type Seller interface {
    Sell() interface{}
}

type RedWineFactory struct {

}

type RedWineProxy struct {
    factory RedWineFactory
    sellCount int
}

func (f *RedWineFactory) Sell() interface{} {
    fmt.Println("This is real role")
    return nil
}

func (p *RedWineProxy) Sell() interface{} {
    obj := p.factory.Sell()
    p.sellCount++
    fmt.Println("The factory selled :", p.sellCount)
    return obj
}

func main() {
    var sell Seller = &RedWineProxy{sellCount: 4}
    sell.Sell()
}
