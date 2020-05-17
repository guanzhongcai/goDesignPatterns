package main

import "design.patterns/BehavioralPatterns/Observer"

func main() {

    subject := Observer.NewSubject()
    r1 := Observer.NewReader("GZC")
    r2 := Observer.NewReader("James")
    r3 := Observer.NewReader("Curry")
    subject.Attach(r1)
    subject.Attach(r2)
    subject.Attach(r3)
    subject.UpdateContext("妹子来了")
    subject.UpdateContext("妹子走了")
}
