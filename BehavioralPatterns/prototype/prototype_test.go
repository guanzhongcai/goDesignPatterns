package prototype

import "testing"

func TestNewManager(t *testing.T) {
    mgr := NewManager()
    t1 := &Type1{name: "type1"}
    mgr.Set("t1", t1)
    t11 := mgr.Get("t1")
    t22 := t11.Clone()
    if t11 == t22 {
        t.Log("type1.clone相等 深拷贝")
    } else {
        t.Log("type1.clone不等 浅拷贝")
    }

    t2 := &Type2{name: "type2"}
    mgr.Set("t2", t2)
    t11 = mgr.Get("t2")
    t22 = t11.Clone()
    if t11 == t22 {
        t.Log("type2.clone相等 深拷贝")
    } else {
        t.Log("type2.clone不等 浅拷贝")
    }
}