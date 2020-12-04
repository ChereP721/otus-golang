package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	runeStr := []rune(str)
	if unicode.IsDigit(runeStr[0]) {
		return "", ErrInvalidString
	}

	var unpackSB unpackerStringBuilder
	var err error

	iRuneStrMax := len(runeStr) - 1
	countContinue := 0 // в текущей версиии алгоритма переменная принимает значения только 0 и 1, но оставил int на случай его усложнения

	for iSymbol, symbol := range runeStr {
		if countContinue > 0 {
			countContinue--
			continue
		}

		if symbol == '\\' {
			if iSymbol >= iRuneStrMax {
				return "", ErrInvalidString
			}

			unpackSB.writeSymbol()

			nextSymbol := runeStr[iSymbol+1]
			if nextSymbol == '\\' || unicode.IsDigit(nextSymbol) {
				unpackSB.setNewSymbol(nextSymbol)
				countContinue++
				continue
			}

			return "", ErrInvalidString
		}

		if unicode.IsDigit(symbol) {
			if unpackSB.repeatSymbol <= 0 {
				return "", ErrInvalidString
			}

			unpackSB.repeatCount, err = strconv.Atoi(string(symbol))
			if err != nil {
				return "", err
			}

			unpackSB.writeSymbol()
			continue
		}

		// не цифра и не \
		unpackSB.writeSymbol()
		unpackSB.setNewSymbol(symbol)
	}

	unpackSB.writeSymbol()

	return unpackSB.String(), nil
}
