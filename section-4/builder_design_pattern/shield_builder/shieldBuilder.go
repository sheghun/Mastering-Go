package shield_builder

import (
	"strings"
)

type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shBuilder struct {
	code string
}

func NewShieldBuilder() *shBuilder {
	return new(shBuilder)
}

func (sh *shBuilder) RaiseFront() *shBuilder {
	sh.code += "F"
	return sh
}

func (sh *shBuilder) RaiseBack() *shBuilder {
	sh.code += "B"
	return sh
}

func (sh *shBuilder) RaiseRight() *shBuilder {
	sh.code += "R"
	return sh
}

func (sh *shBuilder) RaiseLeft() *shBuilder {
	sh.code += "L"
	return sh
}

func (sh *shBuilder) Build() *shield {
	code := sh.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		left:  strings.Contains(code, "L"),
		right: strings.Contains(code, "R"),
	}
}
