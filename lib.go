package gotestlib

func ExportedFunc() int {
	return 7
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
