package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_NUM = 100000
	MIN_NUM = -100000
)

type Flag struct {
	Median            bool
	Mean              bool
	Mode              bool
	StandartDeviation bool
}

func Main() {

}

func Scan() []int {
	res := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if num > MAX_NUM || num < MIN_NUM {
			fmt.Println("only between -100000 and 100000")
			continue
		}
		res = append(res, num)
	}
	return res
}
