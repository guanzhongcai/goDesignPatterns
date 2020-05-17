package Factory

//实际运行类的接口
type Operator interface {
    SetLeft(int)
    SetRight(int)
    Result() int
}

//工厂接口
type OperatorFactory interface {
    Create() Operator
}
