package impl

import "golandIssue0/foo"

var _ foo.Bar = (*bar)(nil)
var _ foo.Bar = (*overStructBar)(nil)
var _ foo.Bar = (*badBar)(nil)

type bar struct {
	foo.BarImpl
}

func (b *bar) DoStuff() error {
	return nil
}

type overStructBar struct {
	foo.BarImplStruct
}

func (b *overStructBar) DoStuff() error {
	return nil
}

type badBar struct {
}

func (b *badBar) DoStuff() error {
	return nil
}

func (b *badBar) duckTypingHelper() {
}
