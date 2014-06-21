    package main

    import (
        "fmt"
    )

    // command interface
    type Commander interface {
        execute()
    }

    // who receive the command
    type Recever struct {
    }

    // receiver do something when the command is received
    func (r *Recever) action() {
        fmt.Println("the Recever action has been token!")
    }

    type Invoker struct {
        Commander
    }

    func (i *Invoker) action() {
        i.Commander.execute()
    }

    // the concrete command
    type ConcreteCommand struct {
        *Recever
    }

    func (c *ConcreteCommand) execute() {
        c.Recever.action()
    }

    func main() {
        receiver := new(Recever) // &Receiver{}
        command := &ConcreteCommand{receiver}

        invoker := Invoker{command}
        invoker.action()
    }
