package thespine

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// theSize represents the size of grouping used by the anagram encode/decode.
const theSize = 3

// ErrInvalidString represents an error for invalid UTF-8 string.
var ErrInvalidString = errors.New("invalid string")

// Decode takes a UTF-8 string as an input and decodes the anagram.
// Error returned in case of an invalid UTF-8 string.
func Decode(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}
	sr := []rune(s)
	l := len(sr)
	if l <= theSize {
		return s, nil
	}

	g := make([][]rune, 0)
	gc := l / theSize
	if l%theSize != 0 {
		gc++
	}
	for i := range gc {
		si := l - (i+1)*theSize
		ei := l - i*theSize
		if si < 0 {
			si = 0
		}
		gs := sr[si:ei]
		g = append(g, gs)
	}

	return runestring(g), nil
}

// Encode takes a UTF-8 string as an input and generates an anagram out of it.
// Error returned in case of an invalid UTF-8 string.
func Encode(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}
	sr := []rune(s)
	l := len(sr)
	if l <= theSize {
		return s, nil
	}

	g := make([][]rune, 0)
	gc := l / theSize
	if l%theSize != 0 {
		gc++
	}
	for i := range gc {
		si := i * theSize
		ei := (i + 1) * theSize
		if ei > l {
			ei = l
		}
		gs := sr[si:ei]
		g = append(g, gs)
	}
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}

	return runestring(g), nil
}

// EncodeText takes a UTF-8 string as an input, splits it by whitespace and runs an anagram for each word.
// Error returned in case of an invalid UTF-8 string.
func EncodeText(s string) (string, error) {
	o := ""
	ws := strings.Split(s, " ")
	for i, w := range ws {
		ew, err := Encode(w)
		if err != nil {
			return "", err
		}
		o += ew
		if i != len(ws)-1 {
			o += " "
		}
	}

	return o, nil
}

// DecodeText takes a UTF-8 string as an input, splits it by whitespace and decodes each anagram word-by-word.
// Error returned in case of an invalid UTF-8 string.
func DecodeText(s string) (string, error) {
	o := ""
	ws := strings.Split(s, " ")
	for i, w := range ws {
		ew, err := Decode(w)
		if err != nil {
			return "", err
		}
		o += ew
		if i < len(ws)-1 {
			o += " "
		}
	}

	return o, nil
}

func runestring(r [][]rune) string {
	var s string
	for _, r := range r {
		s += string(r)
	}

	return s
}
