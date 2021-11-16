package sum

func sum(array []int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

func numbersToSum(numbers ...[]int) []int {
	result := make([]int, len(numbers))
	for i, itemOfNumbers := range numbers {
		result[i] = sum(itemOfNumbers)
	}
	return result
}

func numbersToSumTail(numbers ...[]int) []int {
	result := make([]int, len(numbers))
	for i, itemOfNumbers := range numbers {
		if len(itemOfNumbers) == 0 {
			result[i] = 0
		} else {
			result[i] = sum(itemOfNumbers[1:])
		}
	}
	return result
}
