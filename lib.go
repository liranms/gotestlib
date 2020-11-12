package gotestlib

func ExportedFunc(i int) float64 {
	return float64(i)
}

type SomeStruct struct {
	Foo int
}

type Inyourface interface {
	Bar()
}
