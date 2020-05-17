package Factory

//数据
type OperatorBase struct {
    left, right int
}

//赋值
func (o *OperatorBase) SetLeft(left int) {

    o.left = left
}

func (o *OperatorBase) SetRight(right int) {

    o.right = right
}
