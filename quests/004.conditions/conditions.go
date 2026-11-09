package conditions

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	if age <= 0 || balance < 0 {
		return "INVALID"
	} else if age < 18 || !hasID {
		return "REJECTED"
	} else if isVIP && balance >= 10000 {
		return "VIP_ACCESS"
	} else if balance >= 1000 {
		return "STANDARD_ACCESS"
	} else {
		return "LIMITED_ACCESS"
	}
}

func EvaluateGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score >= 0 && score < 60:
		return "F"
	default:
		return "INVALID"
	}

}
