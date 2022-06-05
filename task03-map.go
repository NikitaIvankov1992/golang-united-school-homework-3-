package homework

import (
	"sort"
)

// function that returns map values sorted in order of increasing keys.
// Input -> {2: "a", 0: "b", 1: "c"}
// Output -> ["b", "c", "a"]
// Input -> {10: "aa", 0: "bb", 500: "cc"}
// Output -> ["bb", "aa", "cc"]

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	output := []string{}

	for k := range keys {
		output = append(output, input[k])
	}
	return output
}
