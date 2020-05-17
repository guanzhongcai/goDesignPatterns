package strategy

import "fmt"

type Context struct {
    Strategy
}

type Strategy interface {
    Do()
}

type Strategy1 struct {
}

type Strategy2 struct {
}

func (s1 *Strategy1) Do() {
    fmt.Println("Strategy1 do")
}

func (s2 *Strategy2) Do() {
    fmt.Println("Strategy2 do")
}
