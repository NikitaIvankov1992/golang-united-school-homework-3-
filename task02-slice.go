package homework

// function that returns the copy of the original slice in reverse order. The type of elements is int64.
// Input -> (1, 2, 5, 15)
// Output -> (15, 5, 2, 1)

func reverse(input []int64) (result []int64) {
	output := []int64{}

	for i := range input {
		output = append(output, input[len(input)-i-1])
	}
	return output
}
