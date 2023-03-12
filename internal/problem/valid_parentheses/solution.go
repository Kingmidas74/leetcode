package valid_parentheses

// Solve for https://leetcode.com/problems/valid-parentheses/
func (s *Solution) Solve(input string) bool {
	if len(input) == 0 {
		return true
	}

	if len(input) == 1 {
		return false
	}

	t := make([]rune, 0)
	for _, r := range input {
		r1 := string(r)
		if r1 == "(" || r1 == "{" || r1 == "[" {
			t = append(t, r)
			continue
		}
		if (r1 == ")" || r1 == "}" || r1 == "]") && len(t) == 0 {
			return false
		}

		if len(t) > 0 {

			l := string(t[len(t)-1])

			if r1 == ")" && l != "(" {
				return false
			}
			if r1 == "]" && l != "[" {
				return false
			}
			if r1 == "}" && l != "{" {
				return false
			}

			t = t[:len(t)-1]
		}
	}

	return len(t) == 0
}
