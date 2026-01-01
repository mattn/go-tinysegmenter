package tinysegmenter

import (
	"reflect"
	"testing"
)

func TestSegment(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"私の名前は中野です", []string{"私", "の", "名前", "は", "中野", "です"}},
		{"2023年12月31日", []string{"2023", "年", "12", "月", "31", "日"}}, // Numbers and kanji should be split
		{"123456", []string{"123456"}}, // Consecutive numbers should not be split
	}

	ts := New()
	for _, tt := range tests {
		result := ts.Segment(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Input: %s, Expected %v, got %v", tt.input, tt.expected, result)
		}
	}
}
