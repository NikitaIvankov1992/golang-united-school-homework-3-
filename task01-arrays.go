package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0

	for _, element := range input {
		sum += element
	}

	return sum / 15
}
