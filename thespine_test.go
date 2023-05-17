package thespine

import (
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	tests := []struct {
		str  string
		name string
		want string
	}{
		{
			str:  "",
			name: "empty string",
			want: "",
		},
		{
			str:  "erecshyrinol",
			name: "the song",
			want: "nolyricshere",
		},
		{
			str:  "nespithe",
			name: "the album",
			want: "thespine",
		},
		{
			str:  "seteernkub",
			name: "the tech",
			want: "kubernetes",
		},
		{
			str:  "ᚬᚩᚡᚣ",
			name: "the runes",
			want: "ᚩᚡᚣᚬ",
		},
		{
			str:  "\xf0\x9f\x9a\x80ketroc",
			name: "the cringe",
			want: "rocket\xf0\x9f\x9a\x80",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Decode(test.str)
			if got != test.want {
				t.Fatalf("decode got: %s\nwant: %s\n", got, test.want)
			}
			got = Encode(test.want)
			if got != test.str {
				t.Fatalf("encode got: %s\nwant: %s\n", got, test.str)
			}
		})
	}
}
