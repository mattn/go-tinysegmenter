package tinysegmenter

import "regexp"

type chartype struct {
	re    *regexp.Regexp
	ctype string
}

type charmap map[string]int

type TinySegmenter struct {
	chartype        []chartype
	preserveList    []string // List of strings that should not be segmented
	preserveTokens  bool     // Flag to preserve programming tokens (URLs, paths, etc.)
	BIAS         int
	BC1          charmap
	BC2          charmap
	BC3          charmap
	BP1          charmap
	BP2          charmap
	BQ1          charmap
	BQ2          charmap
	BQ3          charmap
	BQ4          charmap
	BW1          charmap
	BW2          charmap
	BW3          charmap
	TC1          charmap
	TC2          charmap
	TC3          charmap
	TC4          charmap
	TQ1          charmap
	TQ2          charmap
	TQ3          charmap
	TQ4          charmap
	TW1          charmap
	TW2          charmap
	TW3          charmap
	TW4          charmap
	UC1          charmap
	UC2          charmap
	UC3          charmap
	UC4          charmap
	UC5          charmap
	UC6          charmap
	UP1          charmap
	UP2          charmap
	UP3          charmap
	UQ1          charmap
	UQ2          charmap
	UQ3          charmap
	UW1          charmap
	UW2          charmap
	UW3          charmap
	UW4          charmap
	UW5          charmap
	UW6          charmap
	NN           charmap
}

func New() *TinySegmenter {
	ts := &TinySegmenter{}
	ts.initPatterns()
	ts.initWeights()
	return ts
}

func (ts *TinySegmenter) initPatterns() {
	patterns := map[string]string{
		"[一二三四五六七八九十百千万億兆]": "M",
		"[一-龠々〆ヵヶ]":         "H",
		"[ぁ-ん]":             "I",
		"[ァ-ヴーｱ-ﾝﾞｰ]":       "K",
		"[a-zA-Zａ-ｚＡ-Ｚ]":    "A",
		"[0-9０-９]":          "N",
	}

	for pattern, ctype := range patterns {
		ts.chartype = append(ts.chartype, chartype{
			re:    regexp.MustCompile(pattern),
			ctype: ctype,
		})
	}
}

func (ts *TinySegmenter) initWeights() {
	ts.BIAS = -332
	ts.BC1 = charmap{"HH": 6, "II": 2461, "KH": 406, "OH": -1378}
	ts.BC2 = charmap{"AA": -3267, "AI": 2744, "AN": -878, "HH": -4070, "HM": -1711, "HN": 4012, "HO": 3761, "IA": 1327, "IH": -1184, "II": -1332, "IK": 1721, "IO": 5492, "KI": 3831, "KK": -8741, "MH": -3132, "MK": 3334, "NM": 15000, "OO": -2920}
	ts.BC3 = charmap{"HH": 996, "HI": 626, "HK": -721, "HN": -1307, "HO": -836, "IH": -301, "KK": 2762, "MK": 1079, "MM": 4034, "OA": -1652, "OH": 266}
	ts.BP1 = charmap{"BB": 295, "OB": 304, "OO": -125, "UB": 352}
	ts.BP2 = charmap{"BO": 60, "OO": -1762}
	ts.initBQWeights()
	ts.initBWWeights()
	ts.initTCWeights()
	ts.initTQWeights()
	ts.initTWWeights()
	ts.initUCWeights()
	ts.initUPWeights()
	ts.initUQWeights()
	ts.initUWWeights()
	ts.NN = charmap{"NN": -11097}
}

func (ts *TinySegmenter) ctype(str string) string {
	for _, ct := range ts.chartype {
		if ct.re.MatchString(str) {
			return ct.ctype
		}
	}
	return "O"
}

// SetPreserveList sets a list of strings that should not be segmented
func (ts *TinySegmenter) SetPreserveList(words []string) {
	ts.preserveList = words
}

// SetPreserveTokens enables/disables preservation of programming tokens
func (ts *TinySegmenter) SetPreserveTokens(enable bool) {
	ts.preserveTokens = enable
}

func (ts *TinySegmenter) ts(v int) int {
	return v
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
		ctype = append(ctype, ts.ctype(string(r)))
	}

	seg = append(seg, "E1", "E2", "E3")
	ctype = append(ctype, "O", "O", "O")

	word := seg[3]
	p1, p2, p3 := "U", "U", "U"

	for i := 4; i < len(seg)-3; i++ {
		score := ts.BIAS
		w1, w2, w3, w4, w5, w6 := seg[i-3], seg[i-2], seg[i-1], seg[i], seg[i+1], seg[i+2]
		c1, c2, c3, c4, c5, c6 := ctype[i-3], ctype[i-2], ctype[i-1], ctype[i], ctype[i+1], ctype[i+2]

		score += ts.ts(ts.UP1[p1])
		score += ts.ts(ts.UP2[p2])
		score += ts.ts(ts.UP3[p3])
		score += ts.ts(ts.BP1[p1+p2])
		score += ts.ts(ts.BP2[p2+p3])
		score += ts.ts(ts.UW1[w1])
		score += ts.ts(ts.UW2[w2])
		score += ts.ts(ts.UW3[w3])
		score += ts.ts(ts.UW4[w4])
		score += ts.ts(ts.UW5[w5])
		score += ts.ts(ts.UW6[w6])
		score += ts.ts(ts.BW1[w2+w3])
		score += ts.ts(ts.BW2[w3+w4])
		score += ts.ts(ts.BW3[w4+w5])
		score += ts.ts(ts.TW1[w1+w2+w3])
		score += ts.ts(ts.TW2[w2+w3+w4])
		score += ts.ts(ts.TW3[w3+w4+w5])
		score += ts.ts(ts.TW4[w4+w5+w6])
		score += ts.ts(ts.UC1[c1])
		score += ts.ts(ts.UC2[c2])
		score += ts.ts(ts.UC3[c3])
		score += ts.ts(ts.UC4[c4])
		score += ts.ts(ts.UC5[c5])
		score += ts.ts(ts.UC6[c6])
		score += ts.ts(ts.BC1[c2+c3])
		score += ts.ts(ts.BC2[c3+c4])
		score += ts.ts(ts.BC3[c4+c5])
		score += ts.ts(ts.TC1[c1+c2+c3])
		score += ts.ts(ts.TC2[c2+c3+c4])
		score += ts.ts(ts.TC3[c3+c4+c5])
		score += ts.ts(ts.TC4[c4+c5+c6])
		score += ts.ts(ts.UQ1[p1+c1])
		score += ts.ts(ts.UQ2[p2+c2])
		score += ts.ts(ts.UQ3[p3+c3])
		score += ts.ts(ts.BQ1[p2+c2+c3])
		score += ts.ts(ts.BQ2[p2+c3+c4])
		score += ts.ts(ts.BQ3[p3+c2+c3])
		score += ts.ts(ts.BQ4[p3+c3+c4])
		score += ts.ts(ts.TQ1[p2+c1+c2+c3])
		score += ts.ts(ts.TQ2[p2+c2+c3+c4])
		score += ts.ts(ts.TQ3[p3+c1+c2+c3])
		score += ts.ts(ts.TQ4[p3+c2+c3+c4])
		score += ts.ts(ts.NN[c3+c4])

		p := "O"
		if score > 0 {
			result = append(result, word)
			word = ""
			p = "B"
		}
		p1, p2, p3 = p2, p3, p
		word += seg[i]
	}
	result = append(result, word)
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
	combined := ""
	for i := pos; i < len(segments); i++ {
		combined += segments[i]
		if combined == target {
			return true, i - pos + 1
		}
		if len(combined) >= len(target) {
			break
		}
	}
	return false, 0
}

// segmentWithTokens performs segmentation while preserving programming tokens
func (ts *TinySegmenter) segmentWithTokens(input string) []string {
	tokenPattern := regexp.MustCompile(`[a-zA-Z0-9._/\-:@]+`)
	
	result := []string{}
	lastEnd := 0
	
	matches := tokenPattern.FindAllStringIndex(input, -1)
	for _, match := range matches {
		start, end := match[0], match[1]
		
		// Process text before token
		if start > lastEnd {
			beforeText := input[lastEnd:start]
			segments := ts.segmentOriginal(beforeText)
			result = append(result, ts.mergePreservedWords(segments)...)
		}
		
		// Add token as single segment
		token := input[start:end]
		result = append(result, token)
		lastEnd = end
	}
	
	// Process remaining text
	if lastEnd < len(input) {
		remainingText := input[lastEnd:]
		segments := ts.segmentOriginal(remainingText)
		result = append(result, ts.mergePreservedWords(segments)...)
	}
	
	return result
}
