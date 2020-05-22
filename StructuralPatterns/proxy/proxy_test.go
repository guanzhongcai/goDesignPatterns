package proxy

import "testing"

func TestRealSubject_Do(t *testing.T) {

    var sub Subject
    sub = &Proxy{money: -1000}
    t.Log(sub.Do())
}
