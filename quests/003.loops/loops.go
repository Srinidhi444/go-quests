package loops

func SumEvenNumbers(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := range n + 1 {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum

}

func KeepOnlyConsonants(strs []string) []string {
	res := []string{}

	for _, str := range strs {
		result := ""

		for _, char := range str {
			switch char {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				continue
			default:
				result += string(char)
			}
		}

		if result != "" {
			res = append(res, result)
		}
	}

	return res
}
