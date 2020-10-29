package sum

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0

	if len(numbers) == 0 {
		return 0
	}

	for _, n := range numbers {
		sum += n
	}

	return sum
}
