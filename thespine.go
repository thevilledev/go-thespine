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

	gc := (l + theSize - 1) / theSize
	g := make([][]rune, gc)
	for i := range gc {
		si := l - (i+1)*theSize
		ei := l - i*theSize
		if si < 0 {
			si = 0
		}
		g[i] = sr[si:ei]
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

	gc := (l + theSize - 1) / theSize
	g := make([][]rune, gc)
	for i := range gc {
		si := i * theSize
		ei := (i + 1) * theSize
		if ei > l {
			ei = l
		}
		g[i] = sr[si:ei]
	}
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}

	return runestring(g), nil
}

// EncodeText takes a UTF-8 string as an input, splits it by whitespace and runs an anagram for each word.
// Error returned in case of an invalid UTF-8 string.
func EncodeText(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var builder strings.Builder
	ws := strings.Split(s, " ")
	for i, w := range ws {
		if w == "" {
			continue // Skip empty strings or preserve them, depending on requirements
		}

		ew, err := Encode(w)
		if err != nil {
			return "", err
		}

		builder.WriteString(ew)
		if i != len(ws)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String(), nil
}

// DecodeText takes a UTF-8 string as an input, splits it by whitespace and decodes each anagram word-by-word.
// Error returned in case of an invalid UTF-8 string.
func DecodeText(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var builder strings.Builder
	ws := strings.Split(s, " ")
	for i, w := range ws {
		if w == "" {
			continue // Skip empty strings or preserve them, depending on requirements
		}

		ew, err := Decode(w)
		if err != nil {
			return "", err
		}

		builder.WriteString(ew)
		if i < len(ws)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String(), nil
}

func runestring(r [][]rune) string {
	// Calculate total capacity needed
	totalCap := 0
	for _, runes := range r {
		totalCap += len(runes)
	}

	builder := strings.Builder{}
	builder.Grow(totalCap) // Pre-allocate exact size needed
	for _, runes := range r {
		builder.WriteString(string(runes))
	}

	return builder.String()
}
