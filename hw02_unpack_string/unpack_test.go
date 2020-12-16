package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			input:    "т1е2с3т4 р3у2с1с0к-1и-2х-3 буков",
			expected: "теессстттт рррууск-и--х--- буков",
		},
		{
			input:    "и 0про 1бе 2лов",
			expected: "ипро бе  лов",
		},
		{
			input:    "#2 &3 $2 << >> & |",
			expected: "## &&& $$ << >> & |",
		},
		{
			input:    "\u00043ы2",
			expected: "\u0004\u0004\u0004ыы",
		},
		{
			input:    "👍3",
			expected: "👍👍👍",
		},
		{
			input:    "ȸ1֍2ץ2ש4ק", // RTL
			expected: "ȸ֍֍ץץששששק",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
		{
			input:    `////5///`,
			expected: `///////////`,
		},
		{
			input:    `qwe\\\`,
			expected: ``,
			err:      ErrInvalidString,
		},
		{
			input:    `qw\ne`,
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    `\👍3`,
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    `\34 asd2`,
			expected: "3333 asdd",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
