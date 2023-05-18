package thespine

import (
	"testing"
)

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
				t.Fatalf("decode got: %s\nwant: %s\n", got, test.want)
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
				t.Fatalf("encode got: %s\nwant: %s\n", enc, test.want)
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
