package thespine

import (
	"fmt"
	"log"
	"testing"
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
			o, _ := EncodeText(test.str)
			if o != test.want {
				t.Fatalf("encode got: '%s'\nwant: '%s'\n", o, test.want)
			}
		})
	}
}

func Test_DecodeText(t *testing.T) {
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
			o, _ := DecodeText(test.str)
			if o != test.want {
				t.Fatalf("encode got: '%s'\nwant: '%s'\n", o, test.want)
			}
		})
	}
}
