package observer

type Subject struct {
    obs     [] Observer //群聊
    content string      //更新群聊的信息
}

func NewSubject() *Subject {

    return &Subject{obs: make([]Observer, 0)}
}

func (s *Subject) Attach(o Observer) {
    s.obs = append(s.obs, o)
}

func (s *Subject) Notify() {
    for _, o := range s.obs { //每个人都发一次
        o.Update(s) //更新通知
    }
}

func (s *Subject) UpdateContext(context string) {
    s.content = context
    s.Notify() //更新之后通知
}
