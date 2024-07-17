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
	array := make(map[int]int)
	for _, value := range numbers {
		if val, ok := array[value]; ok {
			array[value] = val + 1
		} else {
			array[value] = 1
		}
	}

	occurrence := array[numbers[0]]
	res := numbers[0]
	for key, value := range array {
		if occurrence < value {
			occurrence = value
			res = key
		}
	}
	return res
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
