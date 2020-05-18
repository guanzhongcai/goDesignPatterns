package adapter

import "testing"

func TestNewAdapter(t *testing.T) {
    adaptee := NewAdaptee() //适配器
    target := NewAdapter(adaptee) //传递进入
    res := target.Request()
    t.Log(res)
}