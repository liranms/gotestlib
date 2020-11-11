package gotestlib

func ExportedFunc() int {
	return 4
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
