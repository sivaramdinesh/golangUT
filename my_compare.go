package hacker
// Compare returns an integer comparing two strings lexicographically.
func Compare(a, b string) int {
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
