package anscombe

func Mean(num []int) float64 {
	sum := 0
	for _, r := range num {
		sum += r
	}
	return float64(sum) / float64(len(num))
}

func Median(num []int) float64 {

}
