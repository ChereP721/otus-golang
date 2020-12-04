package hw02_unpack_string //nolint:golint,stylecheck

import (
	"strings"
)

type unpackerStringBuilder struct {
	strings.Builder
	repeatSymbol rune
	repeatCount  int
}

func (ub *unpackerStringBuilder) writeSymbol() {
	if ub.repeatSymbol > 0 {
		_, _ = ub.WriteString(strings.Repeat(string(ub.repeatSymbol), ub.repeatCount))
		ub.repeatSymbol = 0
		ub.repeatCount = 1
	}
}
func (ub *unpackerStringBuilder) setNewSymbol(symbol rune) {
	ub.repeatSymbol = symbol
	ub.repeatCount = 1
}
