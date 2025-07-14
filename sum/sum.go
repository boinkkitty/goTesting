package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, v := range numbers {
		sums = append(sums, Sum(v))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, v := range numbers {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
