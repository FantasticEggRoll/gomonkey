package abc

type Test struct {
	A int
	B int
}

var FnVal = func() int {
	return 1
}

func (t *Test) Hello() int {
	return t.A
}

func (t Test) Hi() int {
	return t.B
}

func HelloWorld() (error, string) {
	return nil, "hello world"
}
