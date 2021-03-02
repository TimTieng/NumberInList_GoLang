package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// Iterate through the List
	for _, val := range list {
		// Evaluate at the index
		if val == num {
			return true
		}
	}
	// Base Case
	return false
}
