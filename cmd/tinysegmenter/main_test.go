package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/mattn/go-tinysegmenter"
)

func TestDoReader(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single line",
			input: "私の名前は中野です\n",
			want:  "私 の 名前 は 中野 です\n",
		},
		{
			name:  "multiple lines",
			input: "私の名前は中野です\nこれはペンです\n",
			want:  "私 の 名前 は 中野 です\nこれ は ペン です\n",
		},
		{
			name:  "no trailing newline",
			input: "これはペンです",
			want:  "これ は ペン です\n",
		},
		{
			name:  "empty line",
			input: "これはペンです\n\nこれはペンです\n",
			want:  "これ は ペン です\n\nこれ は ペン です\n",
		},
		{
			name:  "digits and punctuation",
			input: "今日は2026年10月8日です。\n",
			want:  "今日 は 2026 年 10月 8 日 です 。\n",
		},
		{
			name:  "spaces are dropped",
			input: "Hello World\n",
			want:  "Hello World\n",
		},
		{
			name:  "full-width space is dropped",
			input: "これは\u3000ペンです\n",
			want:  "これ は ペン です\n",
		},
		{
			name:  "full-width digits are joined",
			input: "１２３円です\n",
			want:  "１２３ 円 です\n",
		},
		{
			name:  "digits separated by space are not joined",
			input: "2026 10です\n",
			want:  "2026 10 です\n",
		},
		{
			name:  "decimals are joined",
			input: "3.14です\n",
			want:  "3.14 です\n",
		},
		{
			name:  "full-width decimals are joined",
			input: "１．５と２．０\n",
			want:  "１．５ と ２．０\n",
		},
		{
			name:  "decimal followed by kanji",
			input: "１．５倍\n",
			want:  "１．５ 倍\n",
		},
		{
			name:  "dotted numbers are joined",
			input: "ver1.2.3です\n",
			want:  "ver 1.2.3 です\n",
		},
		{
			name:  "trailing point is not joined",
			input: "3.です\n",
			want:  "3 . です\n",
		},
		{
			name:  "leading point is not joined",
			input: "と.5\n",
			want:  "と . 5\n",
		},
		{
			name:  "point separated by space is not joined",
			input: "3 .14\n",
			want:  "3 . 14\n",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out strings.Builder
			if err := doReader(tinysegmenter.New(), strings.NewReader(tt.input), &out); err != nil {
				t.Fatalf("doReader() error = %v", err)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("doReader() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDoReaderPreserveTokens(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "url",
			input: "URLはhttps://example.comです\n",
			want:  "URL は https://example.com です\n",
		},
		{
			name:  "e-mail",
			input: "mattn@example.jpに送る\n",
			want:  "mattn@example.jp に 送る\n",
		},
		{
			name:  "spaces are dropped",
			input: "Hello World\n",
			want:  "Hello World\n",
		},
		{
			name:  "full-width decimals are joined",
			input: "１．５倍\n",
			want:  "１．５ 倍\n",
		},
	}

	seg := tinysegmenter.New()
	seg.SetPreserveTokens(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out strings.Builder
			if err := doReader(seg, strings.NewReader(tt.input), &out); err != nil {
				t.Fatalf("doReader() error = %v", err)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("doReader() = %q, want %q", got, tt.want)
			}
		})
	}
}

type errReader struct{ err error }

func (r errReader) Read([]byte) (int, error) { return 0, r.err }

func TestDoReaderReadError(t *testing.T) {
	want := errors.New("read error")
	var out strings.Builder
	if err := doReader(tinysegmenter.New(), errReader{want}, &out); !errors.Is(err, want) {
		t.Fatalf("doReader() error = %v, want %v", err, want)
	}
}
