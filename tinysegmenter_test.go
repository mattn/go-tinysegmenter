package tinysegmenter

import (
	"reflect"
	"testing"
)

func TestSegment(t *testing.T) {
	tests := []struct {
		input     string
		preserved []string
		expected  []string
	}{
		{"私の名前は中野です", nil, []string{"私", "の", "名前", "は", "中野", "です"}},
		{"2023年12月31日", nil, []string{"2023", "年", "12", "月", "31", "日"}}, // Numbers and kanji should be split
		{"123456", nil, []string{"123456"}},                               // Consecutive numbers should not be split
		{"アリババと40人の盗賊", nil, []string{"アリババ", "と", "40", "人", "の", "盗賊"}}, // Katakana and numbers
		{"2025年大阪万博", []string{"万博"}, []string{"2025", "年", "大阪", "万博"}},
	}

	ts := New()
	for _, tt := range tests {
		ts.SetPreserveList(tt.preserved)
		result := ts.Segment(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Input: %s, Expected %v, got %v", tt.input, tt.expected, result)
		}
	}
}

func TestPreserveTokens(t *testing.T) {
	ts := New()
	ts.SetPreserveTokens(true)
	
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"URLはhttps://example.comです",
			[]string{"URL", "は", "https://example.com", "です"},
		},
		{
			"ファイルはfoo/bar.txtです",
			[]string{"ファイル", "は", "foo/bar.txt", "です"},
		},
		{
			"関数func_nameを呼び出す",
			[]string{"関数", "func_name", "を", "呼び出す"},
		},
		{
			"メールuser@example.comを送信",
			[]string{"メール", "user@example.com", "を", "送信"},
		},
		{
			"foo bar",
			[]string{"foo", " ", "bar"},
		},
	}
	
	for _, test := range tests {
		result := ts.Segment(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Input: %s\nExpected: %v\nGot: %v", test.input, test.expected, result)
		}
	}
}
