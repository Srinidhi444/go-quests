package slice

func ProcessScores(scores []int) []int {
	finalScores := []int{}

	// 1. Filter and 2. Normalize
	for _, score := range scores {
		if score < 0 || score > 100 {
			continue
		}

		if score < 40 {
			score = 40
		}

		finalScores = append(finalScores, score)
	}

	// 3. Bonus and 4. Cap
	if len(finalScores) > 5 {
		for i := 0; i < len(finalScores); i++ {
			finalScores[i] += 5

			if finalScores[i] > 100 {
				finalScores[i] = 100
			}
		}
	}

	return finalScores
}
