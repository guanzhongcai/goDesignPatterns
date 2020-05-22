package composite

import "fmt"

type Leaf struct {
    component
}

func NewLeaf() *Leaf { //开辟一个叶子

    return &Leaf{}
}

func (c *Leaf) Print(pre string) {
    fmt.Println(pre, c.name)
}

type Composite struct {
    component
    childs [] Component //叶子集合
}

//创建一个组合结构体
func NewComposite() *Composite {

    return &Composite{childs: make([]Component, 0)}
}

func (c *Composite) AddChild(child Component) {

    child.SetParent(c)
    c.childs = append(c.childs, child) //加入叶子节点
}

//打印显示每一个节点
func (c *Composite) Print(pre string) {

    fmt.Println(pre, c.name)
    pre += "    "
    for _, comp := range c.childs {
        comp.Print(pre)
    }
}
