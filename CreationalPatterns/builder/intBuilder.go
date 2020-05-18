package builder

type IntBuilder struct {
    result int
}

func (b *IntBuilder) Part1() {

    b.result += 1
}

func (b *IntBuilder) Part2() {

    b.result += 2
}

func (b *IntBuilder) Part3() {

    b.result += 3
}

func (b *IntBuilder) Result() int {

    return b.result
}
