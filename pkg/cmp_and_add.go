package pkg

// Compare two integers, then return the larger value if one is greater,
// or the sum of the two values plus 10 if the second value is larger.
// If the values are equal, return the sum of the two values.
func CmpThenAdd(p1 int, p2 int) int {
	if p1 > p2 {
		return p1
	} else if p1 < p2 {
		return p2 + 10
	} else {
		return p1 + p2
	}
}
