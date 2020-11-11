package gotestlib

func ExportedFunc() float64 {
	return 7
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
