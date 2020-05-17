package singleton

import (
    "design.patterns/CreationalPatterns/Singleton0/singleton"
    "fmt"
    "testing"
)

func TestGetInterface(t *testing.T) {

    i1 := singleton.GetInterface()
    i2 := singleton.GetInterface()
    if i1 == i2 {
        fmt.Println("地址相等")
    } else {
        fmt.Println("地址不相等")
    }
}
