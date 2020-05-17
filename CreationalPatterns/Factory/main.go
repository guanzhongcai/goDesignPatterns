package main

import (
    "design.patterns/CreationalPatterns/Factory/Factory"
    "fmt"
)

func main() {

    var factory Factory.OperatorFactory
    factory = Factory.PlusOperatorFactory{}
    op := factory.Create()
    op.SetLeft(20)
    op.SetRight(10)
    fmt.Println(op.Result())

    factory = Factory.SubOperatorFactory{}
    op = factory.Create()
    op.SetLeft(20)
    op.SetRight(10)
    fmt.Println(op.Result())
}
