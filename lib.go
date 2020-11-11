package gotestlib

func ExportedFunc() int {
	return 3
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
