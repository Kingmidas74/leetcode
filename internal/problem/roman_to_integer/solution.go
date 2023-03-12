package roman_to_integer

// Solve for https://leetcode.com/problems/roman-to-integer/
func (s *Solution) Solve(input string) int {
	t := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	r := []rune(input)
	v := 0
	n, ok := t[string(r[len(r)-1])]
	if !ok {
		return 0
	}
	v = n

	for i := len(r) - 2; i >= 0; i-- {
		cv, ok := t[string(r[i])]
		if !ok {
			return 0
		}
		pv, ok := t[string(r[i+1])]
		if !ok {
			return 0
		}

		if cv < pv {
			v = v - cv
		} else {
			v = v + cv
		}
	}

	return v
}
