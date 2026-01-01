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
		{"2023年12月31日", []string{"2", "0", "2", "3", "年", "1", "2", "月", "3", "1", "日"}}, // Original: all numbers split
		{"123456", []string{"1", "2", "3", "4", "5", "6"}}, // Original: consecutive numbers are split
	}

	ts := New()
	for _, tt := range tests {
		result := ts.Segment(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Input: %s, Expected %v, got %v", tt.input, tt.expected, result)
		}
	}
}
