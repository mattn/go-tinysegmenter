package tinysegmenter

func (ts *TinySegmenter) initTQWeights() {
	ts.TQ1 = map[string]int{"BHHH": -227, "BHHI": 316, "BHIH": -132, "BIHH": 60, "BIII": 1595, "BNHH": -744, "BOHH": 225, "BOOO": -908, "OAKK": 482, "OHHH": 281, "OHIH": 249, "OIHI": 200, "OIIH": -68}
	ts.TQ2 = map[string]int{"BIHH": -1401, "BIII": -1033, "BKAK": -543, "BOOO": -5591}
	ts.TQ3 = map[string]int{"BHHH": 478, "BHHM": -1073, "BHIH": 222, "BHII": -504, "BIIH": -116, "BIII": -105, "BMHI": -863, "BMHM": -464, "BOMH": 620, "OHHH": 346, "OHHI": 1729, "OHII": 997, "OHMH": 481, "OIHH": 623, "OIIH": 1344, "OKAK": 2792, "OKHH": 587, "OKKA": 679, "OOHH": 110, "OOII": -685}
	ts.TQ4 = map[string]int{"BHHH": -721, "BHHM": -3604, "BHII": -966, "BIIH": -607, "BIII": -2181, "OAAA": -2763, "OAKK": 180, "OHHH": -294, "OHHI": 2446, "OHHO": 480, "OHIH": -1573, "OIHH": 1935, "OIHI": -493, "OIIH": 626, "OIII": -4007, "OKAK": -8156}
}

func (ts *TinySegmenter) initTWWeights() {
	ts.TW1 = map[string]int{"につい": -4681, "東京都": 2026}
	ts.TW2 = map[string]int{"ある程": -2049, "いった": -1256, "ころが": -2434, "しょう": 3873, "その後": -4430, "だって": -1049, "ていた": 1833, "として": -4657, "ともに": -4517, "もので": 1882, "一気に": -792, "初めて": -1512, "同時に": -8097, "大きな": -1255, "対して": -2721, "社会党": -3216}
	ts.TW3 = map[string]int{"いただ": -1734, "してい": 1314, "として": -4314, "につい": -5483, "にとっ": -5989, "に当た": -6247, "ので,": -727, "ので、": -727, "のもの": -600, "れから": -3752, "十二月": -2287}
	ts.TW4 = map[string]int{"いう.": 8576, "いう。": 8576, "からな": -2348, "してい": 2958, "たが,": 1516, "たが、": 1516, "ている": 1538, "という": 1349, "ました": 5543, "ません": 1097, "ようと": -4258, "よると": 5865}
}

func (ts *TinySegmenter) initUCWeights() {
	ts.UC1 = map[string]int{"A": 484, "K": 93, "M": 645, "O": -505}
	ts.UC2 = map[string]int{"A": 819, "H": 1059, "I": 409, "M": 3987, "N": 5775, "O": 646}
	ts.UC3 = map[string]int{"A": -1370, "I": 2311}
	ts.UC4 = map[string]int{"A": -2643, "H": 1809, "I": -1032, "K": -3450, "M": 3565, "N": 3876, "O": 6646}
	ts.UC5 = map[string]int{"H": 313, "I": -1238, "K": -799, "M": 539, "O": -831}
	ts.UC6 = map[string]int{"H": -506, "I": -253, "K": 87, "M": 247, "O": -387}
}

func (ts *TinySegmenter) initUPWeights() {
	ts.UP1 = map[string]int{"O": -214}
	ts.UP2 = map[string]int{"B": 69, "O": 935}
	ts.UP3 = map[string]int{"B": 189}
}

func (ts *TinySegmenter) initUQWeights() {
	ts.UQ1 = map[string]int{"BH": 21, "BI": -12, "BK": -99, "BN": 142, "BO": -56, "OH": -95, "OI": 477, "OK": 410, "OO": -2422}
	ts.UQ2 = map[string]int{"BH": 216, "BI": 113, "OK": 1759}
	ts.UQ3 = map[string]int{"BA": -479, "BH": 42, "BI": 1913, "BK": -7198, "BM": 3160, "BN": 6427, "BO": 14761, "OI": -827, "ON": -3212}
}
