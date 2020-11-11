package gotestlib

func ExportedFunc() int {
	return 5
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
