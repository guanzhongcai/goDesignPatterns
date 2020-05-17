package Observer

type Observer interface {
    Update(subject *Subject) //更新一个消息
}
