package garbage

var breathLetters = map[string]string{
	"a":  "ａ",
	"b":  "ｂ",
	"c":  "ｃ",
	"d":  "ｄ",
	"e":  "ｅ",
	"f":  "ｆ",
	"g":  "ｇ",
	"h":  "ｈ",
	"i":  "ｉ",
	"j":  "ｊ",
	"k":  "ｋ",
	"l":  "ｌ",
	"m":  "ｍ",
	"n":  "ｎ",
	"o":  "ｏ",
	"p":  "ｐ",
	"q":  "ｑ",
	"r":  "ｒ",
	"s":  "ｓ",
	"t":  "ｔ",
	"u":  "ｕ",
	"v":  "ｖ",
	"w":  "ｗ",
	"x":  "ｘ",
	"y":  "ｙ",
	"z":  "ｚ",
	"A":  "Ａ",
	"B":  "Ｂ",
	"C":  "Ｃ",
	"D":  "Ｄ",
	"E":  "Ｅ",
	"F":  "Ｆ",
	"G":  "Ｇ",
	"H":  "Ｈ",
	"I":  "Ｉ",
	"J":  "Ｊ",
	"K":  "Ｋ",
	"L":  "Ｌ",
	"M":  "Ｍ",
	"N":  "Ｎ",
	"O":  "Ｏ",
	"P":  "Ｐ",
	"Q":  "Ｑ",
	"R":  "Ｒ",
	"S":  "Ｓ",
	"T":  "Ｔ",
	"U":  "Ｕ",
	"V":  "Ｖ",
	"W":  "Ｗ",
	"X":  "Ｘ",
	"Y":  "Ｙ",
	"Z":  "Ｚ",
	"1":  "１",
	"2":  "２",
	"3":  "３",
	"4":  "４",
	"5":  "５",
	"6":  "６",
	"7":  "７",
	"8":  "８",
	"9":  "９",
	"0":  "０",
	" ":  "　",
	"?":  "？",
	"!":  "！",
	",":  "，",
	";":  "；",
	".":  "．",
	":":  "：",
	"-":  "－",
	"_":  "_",
	"/":  "／",
	"\\": "＼",
	"(":  "（",
	")":  "）",
	"&":  "＆",
	"%":  "％",
	"#":  "＃",
	"+":  "＋",
	"*":  "＊",
	"$":  "＄",
	"@":  "＠",
	"=":  "＝",
	"'":  "＇",
	"[":  "［",
	"]":  "］",
	"~":  "～",
	"^":  "＾",
	"{":  "｛",
	"}":  "｝",
	"|":  "｜",
	"<":  "＜",
	">":  "＞",
	"°":  "ﾟ",
	"£":  "￡",
	"è":  "ｅ＇",
	"é":  "ｅ＇",
	"à":  "ａ＇",
	"ò":  "ｏ＇",
	"ì":  "ｉ＇",
	"ù":  "ｕ＇",
}

// Breath takes a string and returns its breathe representation
func Breath(s string) (r string) {
	for _, c := range s {
		nc := ""
		nc = string(c)
		if q, ok := breathLetters[string(c)]; ok {
			nc = q
		}

		r = r + nc
	}

	return
}
