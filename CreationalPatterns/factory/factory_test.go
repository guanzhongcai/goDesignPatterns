package factory

import (
    "fmt"
    "testing"
)

func TestFactory(t *testing.T) {

    var factory OperatorFactory
    factory = PlusOperatorFactory{}
    op := factory.Create()
    op.SetLeft(20)
    op.SetRight(10)
    fmt.Println(op.Result())

    factory = SubOperatorFactory{}
    op = factory.Create()
    op.SetLeft(20)
    op.SetRight(10)
    fmt.Println(op.Result())
}
