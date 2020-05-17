package singleton

import (
    "sync"
)

type Single struct {
    data int
}

var singleton *Single
var once sync.Once //内核信号，时时刻刻只能运行一个

func GetInterface() *Single {

    once.Do(func() {
        singleton = &Single{data: 100}
    })

    return singleton
}
