package Factory

//操作的抽象
type PlusOperatorFactory struct {

}
type PlusOperator struct {
    *OperatorBase
}

//实际运行
func (o PlusOperator) Result() int {
    return o.left + o.right
}

func (PlusOperatorFactory) Create() Operator {

    return &PlusOperator{&OperatorBase{}}
}
