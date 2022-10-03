package ognoc

// Detect detects the complicated of cipher by easy scoring mechanism
func Detect(result string) string {
	a := scoreByLen(result)
	b := scoreByLetter(result)
	c := scoreBySpechar(result)
	d := scoreByNumber(result)
	e := reward(result)
	score := a + b + c + d + e

	difficult := ""
	switch {
	case score >= 90 && score <= 100:
		difficult = "VERY_SECURE"
	case score >= 80 && score < 90:
		difficult = "SECURE"
	case score >= 70 && score < 80:
		difficult = "VERY_STRONG"
	case score >= 60 && score < 70:
		difficult = "STRONG"
	case score >= 50 && score < 60:
		difficult = "AVERAGE"
	case score >= 25 && score < 50:
		difficult = "WEAK"
	case score >= 0 && score < 25:
		difficult = "VERY_WEAK"
	}

	return difficult
}

func scoreByLen(result string) (score int) {
	n := len(result)
	switch {
	case n <= 4:
		score = 5
	case n > 5 && n <= 7:
		score = 10
	case n >= 8:
		score = 25
	}

	return
}

func scoreByLetter(result string) (score int) {
	len1, len2 := calLetter(result)
	n := len1 + len2
	if n <= 0 {
		score = 0
	} else if len1 == 0 || len2 == 0 {
		score = 10
	} else {
		score = 20
	}

	return
}

func scoreBySpechar(result string) (score int) {
	n := calChar(result)
	switch {
	case n == 0:
		score = 0
	case n == 1:
		score = 10
	case n > 1:
		score = 25
	}

	return
}

func scoreByNumber(result string) (score int) {
	n := calNumber(result)
	switch {
	case n == 0:
		score = 0
	case n == 1:
		score = 10
	case n > 1:
		score = 20
	}

	return
}

func reward(result string) (score int) {
	l1, l2 := calLetter(result)
	l := l1 + l2
	num := calNumber(result)
	ch := calChar(result)

	if l != 0 && num != 0 {
		score = 2
	} else if l != 0 && num != 0 && ch != 0 {
		score = 3
	} else if l1 != 0 && l2 != 0 && num != 0 && ch != 0 {
		score = 5
	}

	return
}

func calNumber(result string) int {
	num := 0
	for _, v := range result {
		if v >= '0' && v <= '9' {
			num++
		}
	}

	return num
}

func calChar(result string) int {
	num := 0
	for _, v1 := range result {
		for _, v2 := range Spechar {
			if v1 == v2 {
				num++
			}
		}
	}

	return num
}

func calLetter(result string) (int, int) {
	len1, len2 := 0, 0
	for _, v := range result {
		if v >= 'a' && v <= 'z' {
			len1++
		}

		if v >= 'A' && v <= 'Z' {
			len2++
		}
	}

	return len1, len2
}
