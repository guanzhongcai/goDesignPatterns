package builder

import "testing"

func TestNewDirector(t *testing.T) {

    //b := StringBuilder{}
    b := IntBuilder{}
    d := NewDirector(&b)
    d.MakeData()
    t.Log(b.Result())
}
