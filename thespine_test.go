package thespine

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"unicode/utf8"
)

func ExampleDecode() {
	str := "nespithe"
	o, err := Decode(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
	// Output: thespine
}

func ExampleDecodeText() {
	str := "nespithe erecshyrinol"
	o, err := DecodeText(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
	// Output: thespine nolyricshere
}

func ExampleEncode() {
	str := "nolyricshere"
	o, err := EncodeText(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
	// Output: erecshyrinol
}

func ExampleEncodeText() {
	str := "nolyricshere thespine"
	o, err := EncodeText(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)
	// Output: erecshyrinol nespithe
}

func Test_Decode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str     string
		name    string
		want    string
		wantErr bool
	}{
		{
			str:     "",
			name:    "empty string",
			want:    "",
			wantErr: false,
		},
		{
			str:     "erecshyrinol",
			name:    "the song",
			want:    "nolyricshere",
			wantErr: false,
		},
		{
			str:     "nespithe",
			name:    "the album",
			want:    "thespine",
			wantErr: false,
		},
		{
			str:     "seteernkub",
			name:    "the tech",
			want:    "kubernetes",
			wantErr: false,
		},
		{
			str:     "ᚬᚩᚡᚣ",
			name:    "the runes",
			want:    "ᚩᚡᚣᚬ",
			wantErr: false,
		},
		{
			str:     "\xf2",
			name:    "the invalid string",
			want:    "",
			wantErr: true,
		},
		{
			str:     "\xf0\x9f\x9a\x80ketroc",
			name:    "the cringe",
			want:    "rocket\xf0\x9f\x9a\x80",
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := Decode(test.str)
			if test.wantErr && err == nil {
				t.Fatalf("decode err wanted but got none")
			}
			if got != test.want {
				t.Fatalf("decode got: '%s'\nwant: '%s'\n", got, test.want)
			}
		})
	}
}

func Test_Encode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str     string
		name    string
		want    string
		wantErr bool
	}{
		{
			str:     "",
			name:    "empty string",
			want:    "",
			wantErr: false,
		},
		{
			str:     "nolyricshere",
			name:    "the song",
			want:    "erecshyrinol",
			wantErr: false,
		},
		{
			str:     "\xf2",
			name:    "the invalid string",
			want:    "",
			wantErr: true,
		},
		{
			str:     "rocket\xf0\x9f\x9a\x80",
			name:    "the cringe",
			want:    "\xf0\x9f\x9a\x80ketroc",
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			enc, err := Encode(test.str)
			if test.wantErr && err == nil {
				t.Fatalf("encode err wanted but got none")
			}
			if enc != test.want {
				t.Fatalf("encode got: '%s'\nwant: '%s'\n", enc, test.want)
			}
		})
	}
}

func FuzzEncode(f *testing.F) {
	tcs := []string{"rocket\xf0\x9f\x9a\x80", "abba acdc", "Hello, 世界!"}
	for _, tc := range tcs {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, og string) {
		ec, _ := Encode(og)
		if ec != "" {
			dc, _ := Decode(ec)
			ogr := []rune(og)
			dcr := []rune(dc)
			if og != dc {
				t.Errorf("Before: '%q', after: '%q'", og, dc)
			}
			if len(ogr) != len(dcr) {
				t.Errorf("Length before: %d, after: %d", len(og), len(dc))
			}
		}
	})
}

func Test_EncodeText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str     string
		name    string
		want    string
		wantErr bool
	}{
		{
			str:     "meh noob",
			name:    "the noob",
			want:    "meh bnoo",
			wantErr: false,
		},
		{
			str:     "The Putrefying Road In The 19th Extremity (...Somewhere Inside The Bowels Of The Endlessness...)",
			name:    "the noob2",
			want:    "The gyinrefPut dRoa In The h19t ityremExt ehermew.So(.. ideIns The elsBow Of The ..)ss.snelesEnd",
			wantErr: false,
		},
		{
			str:     "\xf2",
			name:    "the invalid string",
			want:    "",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			o, _ := EncodeText(test.str)
			if o != test.want {
				t.Fatalf("encode got: '%s'\nwant: '%s'\n", o, test.want)
			}
		})
	}
}

func Test_DecodeText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str     string
		name    string
		want    string
		wantErr bool
	}{
		{
			str:     "The gyinrefPut dRoa In The h19t ityremExt ehermew.So(.. ideIns The elsBow Of The ..)ss.snelesEnd",
			name:    "the puzzle",
			want:    "The Putrefying Road In The 19th Extremity (...Somewhere Inside The Bowels Of The Endlessness...)",
			wantErr: false,
		},
		{
			str:     "\xf2",
			name:    "the invalid string",
			want:    "",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			o, _ := DecodeText(test.str)
			if o != test.want {
				t.Fatalf("encode got: '%s'\nwant: '%s'\n", o, test.want)
			}
		})
	}
}

func FuzzEncodeDecodeComprehensive(f *testing.F) {
	// Add seed corpus
	seeds := []string{
		"",                        // Empty string
		"a",                       // Single char
		"ab",                      // Two chars
		"abc",                     // Three chars (group size)
		"abcd",                    // More than group size
		"Hello, 世界!",              // Mixed ASCII and Unicode
		"🌍🌎🌏",                     // Only emojis
		"     ",                   // Multiple spaces
		"a\nb\tc",                 // Special whitespace
		"a  b    c",               // Multiple consecutive spaces
		strings.Repeat("a", 1000), // Long string
		"ᚠᛇᚻ᛫ᛒᛦᚦ",                 // Runes
		"\u200B\u200C\u200D",      // Zero-width characters
		"a\u0300\u0301b\u0302c",   // Combining diacritical marks
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Skip invalid UTF-8
		if !utf8.ValidString(input) {
			return
		}

		// Test 1: Encode->Decode roundtrip
		encoded, err := Encode(input)
		if err != nil {
			// Some inputs might legitimately fail to encode
			return
		}
		decoded, err := Decode(encoded)
		if err != nil {
			t.Errorf("Failed to decode encoded string: %v", err)

			return
		}
		if decoded != input {
			t.Errorf("Roundtrip failed: input=%q, got=%q", input, decoded)
		}

		// Test 2: Check encoded string properties
		inputRunes := []rune(input)
		encodedRunes := []rune(encoded)
		if len(inputRunes) != len(encodedRunes) {
			t.Errorf("Length mismatch: input=%d, encoded=%d", len(inputRunes), len(encodedRunes))
		}

		// Test 3: Multiple encode/decode cycles
		current := input
		for i := range 3 {
			encoded, err := Encode(current)
			if err != nil {
				t.Errorf("Failed at cycle %d: %v", i, err)

				return
			}
			decoded, err := Decode(encoded)
			if err != nil {
				t.Errorf("Failed at cycle %d: %v", i, err)

				return
			}
			if decoded != current {
				t.Errorf("Cycle %d failed: expected=%q, got=%q", i, current, decoded)
			}
			current = decoded
		}
	})
}

func FuzzEncodeDecodeText(f *testing.F) {
	seeds := []string{
		"",
		"hello world",
		"  spaced  words  ",
		"one two three four",
		"Hello,\nWorld!",
		"Tab\there",
		"Mixed 世界 Unicode",
		"🌍 Earth 🌎 Globe 🌏",
		strings.Repeat("word ", 100),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Skip invalid UTF-8
		if !utf8.ValidString(input) {
			return
		}

		// Test 1: EncodeText->DecodeText roundtrip
		encoded, err := EncodeText(input)
		if err != nil {
			return
		}
		decoded, err := DecodeText(encoded)
		if err != nil {
			t.Errorf("Failed to decode encoded text: %v", err)

			return
		}

		// Normalize spaces for comparison since that's part of the spec
		normalizeSpaces := func(s string) string {
			return strings.Join(strings.Fields(s), " ")
		}

		normalizedInput := normalizeSpaces(input)
		normalizedDecoded := normalizeSpaces(decoded)

		if normalizedDecoded != normalizedInput {
			t.Errorf("Roundtrip failed:\ninput=%q\ngot=%q", normalizedInput, normalizedDecoded)
		}

		// Test 2: Check word boundaries are preserved
		inputWords := strings.Fields(input)
		encodedWords := strings.Fields(encoded)
		if len(inputWords) != len(encodedWords) {
			t.Errorf("Word count mismatch: input=%d, encoded=%d", len(inputWords), len(encodedWords))
		}
	})
}

// Add benchmark tests.
func BenchmarkEncode(b *testing.B) {
	inputs := []struct {
		name string
		str  string
	}{
		{"small", "hello"},
		{"medium", strings.Repeat("hello", 100)},
		{"large", strings.Repeat("hello", 1000)},
		{"unicode", "Hello, 世界! 🌍"},
	}

	for _, input := range inputs {
		b.Run(input.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Encode(input.str)
			}
		})
	}
}
