package prototype

//原型对象需要实现的接口
type Cloneable interface {
    Clone() Cloneable
}

//原型对象的类
type Manager struct {
    prototypes map[string]Cloneable
}

func NewManager() *Manager {

    return &Manager{prototypes: make(map[string]Cloneable)}
}

func (p *Manager) Get(name string) Cloneable {

    return p.prototypes[name]
}

func (p *Manager) Set(name string, prototype Cloneable) {

    p.prototypes[name] = prototype
}

type Type1 struct {
    name string
}

func (t *Type1) Clone() Cloneable {

    tc := *t
    return &tc //浅拷贝
}

type Type2 struct {
    name string
}

func (t *Type2) Clone() Cloneable {

    return t //深拷贝
}
