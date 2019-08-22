package foo

type Bar interface {
	BarImpl
	DoStuff() error
}

type BarImpl interface {
	duckTypingHelper()
}

type BarImplStruct struct {
}

func (BarImplStruct) duckTypingHelper() {
}


