package main

import "design.patterns/BehavioralPatterns/Observer/observer"

func main() {

    subject := observer.NewSubject()
    r1 := observer.NewReader("nick1")
    r2 := observer.NewReader("nick2")
    r3 := observer.NewReader("nick3")
    subject.Attach(r1)
    subject.Attach(r2)
    subject.Attach(r3)
    subject.UpdateContext("妹子来了")
    subject.UpdateContext("妹子走了")
}
