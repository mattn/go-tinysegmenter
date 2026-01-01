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
