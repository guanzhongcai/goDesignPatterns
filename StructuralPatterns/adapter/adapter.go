package adapter

//适配的目标接口
type Target interface {
    Request() string
}

//被适配
type Adaptee interface {
    SpecificRequest() string
}

//适配器
type Adapter struct {
    Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
    return &Adapter{adaptee}
}

//接口包装
func (a *Adapter) Request() string {

    return a.SpecificRequest()
}

//载体，被适配的目标类
type adapteeImplement struct {
}

func (a *adapteeImplement) SpecificRequest() string {

    return "SpecificRequest"
}

func NewAdaptee() Adaptee {
    return &adapteeImplement{}
}
