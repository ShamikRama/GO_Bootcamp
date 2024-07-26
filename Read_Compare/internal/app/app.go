package app

import (
	"Day00/internal/anscombe"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_NUM = 100000
	MIN_NUM = -100000
)

type Flags struct {
	Median            bool
	Mean              bool
	Mode              bool
	StandartDeviation bool
}

func Main() {
	var flg Flags
	Parsing(&flg)
	num := Scan()
	if len(num) > 0 {
		Stat(num, flg)
	}

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

func Parsing(flags *Flags) {
	flag.BoolVar(&flags.Median, "median", false, "use median")
	flag.BoolVar(&flags.Mean, "mean", false, "use mean")
	flag.BoolVar(&flags.Mode, "mode", false, "use mode")
	flag.BoolVar(&flags.StandartDeviation, "sd", false, "use sd")
	flag.Parse()
}

func Stat(num []int, flags Flags) {
	if flags.Mean {
		fmt.Printf("%.2f\n ", anscombe.Mean(num))
	}
	if flags.Median {
		fmt.Printf("%.2f\n ", anscombe.Median(num))
	}
	if flags.StandartDeviation {
		fmt.Printf("%.2f\n ", anscombe.StandartDeviation(num))
	}
	if flags.Mode {
		fmt.Print(anscombe.Mode(num))
	}
}
