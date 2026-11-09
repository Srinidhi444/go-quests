package strings

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	byteLength := len(s)
	runeCount := len([]rune(s))
	return TextStats{ByteLength: byteLength, RuneCount: runeCount}
}

func RuneFrequencies(s string) map[rune]int {
	mp := make(map[rune]int)
	for _, r := range s {
		mp[r]++
	}
	return mp
}

func FirstRunePosition(s string, target rune) int {
	for i, value := range s {
		if value == target {
			return i
		}
	}
	return -1
}

func ExtractRunes(s string) []rune {
	res := []rune{}
	for _, r := range s {
		res = append(res, r)
	}
	return res
}

func HasOnlyASCII(s string) bool {
	// TODO: implement
	// Read README.md for the instructions
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
