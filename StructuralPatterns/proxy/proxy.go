package proxy

type Subject interface {
    Do() string //实际业务，业务系统，检查业务是否欠费，检查密码是否正确
}

type RealSubject struct {
}

func (sb RealSubject) Do() string {

    return "执行扣费"
}

type Proxy struct {
    real  RealSubject
    money int
}

func (p Proxy) Do() string {
    if p.money > 0 {
        return p.real.Do()
    } else {
        return "请充值"
    }
}
