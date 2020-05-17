package factory

//操作的抽象
type SubOperatorFactory struct {
}

type SubOperator struct {
    *OperatorBase
}

//实际运行
func (o *SubOperator) Result() int {
    return o.left - o.right
}

func (SubOperatorFactory) Create() Operator {

    return &SubOperator{&OperatorBase{}}
}
