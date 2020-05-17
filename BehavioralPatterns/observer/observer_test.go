package observer

import "testing"

func TestNewSubject(t *testing.T) {
    subject := NewSubject()
    r1 := NewReader("nick1")
    r2 := NewReader("nick2")
    r3 := NewReader("nick3")
    subject.Attach(r1)
    subject.Attach(r2)
    subject.Attach(r3)
    subject.UpdateContext("妹子来了")
    subject.UpdateContext("妹子走了")
}
