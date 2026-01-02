package tinysegmenter

import "strings"

type charmap map[string]int

type TinySegmenter struct {
	preserveList   []string // List of strings that should not be segmented
	preserveTokens bool     // Flag to preserve programming tokens (URLs, paths, etc.)
	_BIAS          int
	_BC1           charmap
	_BC2           charmap
	_BC3           charmap
	_BP1           charmap
	_BP2           charmap
	_BQ1           charmap
	_BQ2           charmap
	_BQ3           charmap
	_BQ4           charmap
	_BW1           charmap
	_BW2           charmap
	_BW3           charmap
	_TC1           charmap
	_TC2           charmap
	_TC3           charmap
	_TC4           charmap
	_TQ1           charmap
	_TQ2           charmap
	_TQ3           charmap
	_TQ4           charmap
	_TW1           charmap
	_TW2           charmap
	_TW3           charmap
	_TW4           charmap
	_UC1           charmap
	_UC2           charmap
	_UC3           charmap
	_UC4           charmap
	_UC5           charmap
	_UC6           charmap
	_UP1           charmap
	_UP2           charmap
	_UP3           charmap
	_UQ1           charmap
	_UQ2           charmap
	_UQ3           charmap
	_UW1           charmap
	_UW2           charmap
	_UW3           charmap
	_UW4           charmap
	_UW5           charmap
	_UW6           charmap
	_NN            charmap
}

func New() *TinySegmenter {
	ts := &TinySegmenter{}
	ts.initWeights()
	return ts
}

func (ts *TinySegmenter) isTokenRune(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == '.' || r == '_' || r == '/' || r == '\\' || r == '-' || r == ':' || r == '@'
}

func (ts *TinySegmenter) initWeights() {
	ts._BIAS = -332
	ts._BC1 = charmap{"HH": 6, "II": 2461, "KH": 406, "OH": -1378}
	ts._BC2 = charmap{"AA": -3267, "AI": 2744, "AN": -878, "HH": -4070, "HM": -1711, "HN": 4012, "HO": 3761, "IA": 1327, "IH": -1184, "II": -1332, "IK": 1721, "IO": 5492, "KI": 3831, "KK": -8741, "MH": -3132, "MK": 3334, "NM": 15000, "OO": -2920}
	ts._BC3 = charmap{"HH": 996, "HI": 626, "HK": -721, "HN": -1307, "HO": -836, "IH": -301, "KK": 2762, "MK": 1079, "MM": 4034, "OA": -1652, "OH": 266}
	ts._BP1 = charmap{"BB": 295, "OB": 304, "OO": -125, "UB": 352}
	ts._BP2 = charmap{"BO": 60, "OO": -1762}
	ts.initBQWeights()
	ts.initBWWeights()
	ts.initTCWeights()
	ts.initTQWeights()
	ts.initTWWeights()
	ts.initUCWeights()
	ts.initUPWeights()
	ts.initUQWeights()
	ts.initUWWeights()
	ts._NN = charmap{"NN": -11097}
}

func (ts *TinySegmenter) ctypeRune(r rune) string {
	switch {
	case (r >= 0x4E00 && r <= 0x9FA0) || r == 0x3005 || r == 0x3006 || r == 0x30F5 || r == 0x30F6:
		return "H"
	case r >= 0x3041 && r <= 0x3093:
		return "I"
	case (r >= 0x30A1 && r <= 0x30F4) || r == 0x30FC || (r >= 0xFF71 && r <= 0xFF9D) || r == 0xFF9E || r == 0xFF70:
		return "K"
	case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= 0xFF41 && r <= 0xFF5A) || (r >= 0xFF21 && r <= 0xFF3A):
		return "A"
	case (r >= '0' && r <= '9') || (r >= 0xFF10 && r <= 0xFF19):
		return "N"
	case r == '一' || r == '二' || r == '三' || r == '四' || r == '五' || r == '六' || r == '七' || r == '八' || r == '九' || r == '十' || r == '百' || r == '千' || r == '万' || r == '億' || r == '兆':
		return "M"
	default:
		return "O"
	}
}

// SetPreserveList sets a list of strings that should not be segmented
func (ts *TinySegmenter) SetPreserveList(words []string) {
	ts.preserveList = words
}

// SetPreserveTokens enables/disables preservation of programming tokens
func (ts *TinySegmenter) SetPreserveTokens(enable bool) {
	ts.preserveTokens = enable
}

func (ts *TinySegmenter) Segment(input string) []string {
	if input == "" {
		return []string{}
	}

	// Extract tokens if preserve flag is enabled
	if ts.preserveTokens {
		return ts.segmentWithTokens(input)
	}

	// Perform normal segmentation
	segments := ts.segmentOriginal(input)

	// Merge preserved words
	return ts.mergePreservedWords(segments)
}

// segmentOriginal performs the original segmentation logic
func (ts *TinySegmenter) segmentOriginal(input string) []string {
	result := []string{}
	seg := []string{"B3", "B2", "B1"}
	ctype := []string{"O", "O", "O"}

	runes := []rune(input)
	for _, r := range runes {
		seg = append(seg, string(r))
		ctype = append(ctype, ts.ctypeRune(r))
	}

	seg = append(seg, "E1", "E2", "E3")
	ctype = append(ctype, "O", "O", "O")

	var word strings.Builder
	word.WriteString(seg[3])
	p1, p2, p3 := "U", "U", "U"

	for i := 4; i < len(seg)-3; i++ {
		score := ts._BIAS
		w1, w2, w3, w4, w5, w6 := seg[i-3], seg[i-2], seg[i-1], seg[i], seg[i+1], seg[i+2]
		c1, c2, c3, c4, c5, c6 := ctype[i-3], ctype[i-2], ctype[i-1], ctype[i], ctype[i+1], ctype[i+2]

		score += ts._UP1[p1]
		score += ts._UP2[p2]
		score += ts._UP3[p3]
		score += ts._BP1[p1+p2]
		score += ts._BP2[p2+p3]
		score += ts._UW1[w1]
		score += ts._UW2[w2]
		score += ts._UW3[w3]
		score += ts._UW4[w4]
		score += ts._UW5[w5]
		score += ts._UW6[w6]
		score += ts._BW1[w2+w3]
		score += ts._BW2[w3+w4]
		score += ts._BW3[w4+w5]
		score += ts._TW1[w1+w2+w3]
		score += ts._TW2[w2+w3+w4]
		score += ts._TW3[w3+w4+w5]
		score += ts._TW4[w4+w5+w6]
		score += ts._UC1[c1]
		score += ts._UC2[c2]
		score += ts._UC3[c3]
		score += ts._UC4[c4]
		score += ts._UC5[c5]
		score += ts._UC6[c6]
		score += ts._BC1[c2+c3]
		score += ts._BC2[c3+c4]
		score += ts._BC3[c4+c5]
		score += ts._TC1[c1+c2+c3]
		score += ts._TC2[c2+c3+c4]
		score += ts._TC3[c3+c4+c5]
		score += ts._TC4[c4+c5+c6]
		score += ts._UQ1[p1+c1]
		score += ts._UQ2[p2+c2]
		score += ts._UQ3[p3+c3]
		score += ts._BQ1[p2+c2+c3]
		score += ts._BQ2[p2+c3+c4]
		score += ts._BQ3[p3+c2+c3]
		score += ts._BQ4[p3+c3+c4]
		score += ts._TQ1[p2+c1+c2+c3]
		score += ts._TQ2[p2+c2+c3+c4]
		score += ts._TQ3[p3+c1+c2+c3]
		score += ts._TQ4[p3+c2+c3+c4]
		score += ts._NN[c3+c4]

		p := "O"
		if score > 0 {
			result = append(result, word.String())
			word.Reset()
			p = "B"
		}
		p1, p2, p3 = p2, p3, p
		word.WriteString(seg[i])
	}
	result = append(result, word.String())
	return result
}

// mergePreservedWords merges segments that form preserved words
func (ts *TinySegmenter) mergePreservedWords(segments []string) []string {
	for _, preserve := range ts.preserveList {
		segments = ts.mergeIfMatches(segments, preserve)
	}
	return segments
}

// mergeIfMatches merges consecutive segments if they form the target word
func (ts *TinySegmenter) mergeIfMatches(segments []string, target string) []string {
	result := []string{}
	i := 0

	for i < len(segments) {
		// Try to match target starting from position i
		matched, length := ts.tryMatch(segments, i, target)
		if matched {
			result = append(result, target)
			i += length
		} else {
			result = append(result, segments[i])
			i++
		}
	}
	return result
}

// tryMatch checks if consecutive segments starting at pos form the target
func (ts *TinySegmenter) tryMatch(segments []string, pos int, target string) (bool, int) {
	var combined strings.Builder
	for i := pos; i < len(segments); i++ {
		combined.WriteString(segments[i])
		if combined.String() == target {
			return true, i - pos + 1
		}
		if combined.Len() >= len(target) {
			break
		}
	}
	return false, 0
}

// segmentWithTokens performs segmentation while preserving programming tokens
func (ts *TinySegmenter) segmentWithTokens(input string) []string {
	result := []string{}
	lastEnd := 0

	runes := []rune(input)
	i := 0
	for i < len(runes) {
		if ts.isTokenRune(runes[i]) {
			start := i
			for i < len(runes) && ts.isTokenRune(runes[i]) {
				i++
			}
			end := i

			// Process text before token
			if start > lastEnd {
				beforeText := string(runes[lastEnd:start])
				segments := ts.segmentOriginal(beforeText)
				result = append(result, ts.mergePreservedWords(segments)...)
			}

			// Add token as single segment
			token := string(runes[start:end])
			result = append(result, token)
			lastEnd = end
		} else {
			i++
		}
	}

	// Process remaining text
	if lastEnd < len(runes) {
		remainingText := string(runes[lastEnd:])
		segments := ts.segmentOriginal(remainingText)
		result = append(result, ts.mergePreservedWords(segments)...)
	}

	return result
}
