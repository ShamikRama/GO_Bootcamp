package anscombe

import (
	"math"
	"sort"
)

func Mean(num []int) float64 {
	sum := 0
	res := 0.0
	for _, r := range num {
		sum += r
	}
	res = float64(sum) / float64(len(num))
	return res
}

func Median(num []int) float64 {
	res := 0.0
	sort.Ints(num)
	if len(num)%2 != 0 {
		res = float64(num[len(num)/2])
	} else {
		res = (float64(num[len(num)/2]) + float64(num[len(num)/2-1])) / 2
	}
	return res
}

func Mode(numbers []int) int {
	mode := numbers[0]
	max_count := 0
	for _, val1 := range numbers {
		count := 0
		for _, val2 := range numbers {
			if val1 == val2 {
				count++
			}
		}
		if count > max_count {
			max_count = count
			mode = val1
		}
	}
	return mode
}

func StandartDeviation(numbers []int) float64 {
	res := 0.0
	res = math.Sqrt(variance(numbers))
	return res
}

func variance(numbers []int) float64 {
	res := 0.0
	var total float64
	for _, value := range numbers {
		total += math.Pow(float64(value)-Mean(numbers), 2)
	}
	res = total / float64(len(numbers))
	return res
}
