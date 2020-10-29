package nums

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	l := len(list)

	if l == 0 {
		return false
	}

	for i := 0; i <= l/2; i++ {
		if list[i] == num || list[l-i-1] == num {
			return true
		}
	}

	/*for _, n := range list {
		if n == num {
			return true
		}
	}*/

	return false
}
