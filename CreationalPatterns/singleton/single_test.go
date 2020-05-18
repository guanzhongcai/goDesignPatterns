package singleton

import (
    "fmt"
    "testing"
)

func TestGetInterface(t *testing.T) {

    i1 := GetInterface()
    i2 := GetInterface()
    if i1 == i2 {
        fmt.Println("地址相等")
    } else {
        fmt.Println("地址不相等")
    }
}
