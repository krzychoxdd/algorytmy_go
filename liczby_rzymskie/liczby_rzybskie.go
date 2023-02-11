package main

import (
	"fmt"
)

func main() {
	fmt.Print(romanToInt("IV"))
}

func romanToInt(s string) int {
	romanNums := make(map[string]int)
	romanNums["I"] = 1
	romanNums["V"] = 5
	romanNums["X"] = 10
	romanNums["L"] = 50
	romanNums["C"] = 100
	romanNums["D"] = 500
	romanNums["M"] = 1000

	numsRoman := make(map[int]string)
	numsRoman[1] = "I"
	numsRoman[5] = "V"
	numsRoman[10] = "X"
	numsRoman[50] = "L"
	numsRoman[100] = "C"
	numsRoman[500] = "D"
	numsRoman[1000] = "M"

	// MM + XL + VI
	// 1000 + 1000 + (50 - 10) + 5 + 1
	// 2046

	var sum int = 0

	for k, val := range s {
		l := string(val)
		nL := string(s[k+1])

		num := romanNums[l]
		nextNum := romanNums[nL]
		if num < nextNum {
			sum = sum + (nextNum - num)
			fmt.Print("Presum:")
			fmt.Print(sum)
			fmt.Print("---")
		} else {
			sum = sum + num
		}
	}

	fmt.Println(sum)

	return 0
}
