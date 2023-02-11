package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	fmt.Print("Ilość liczb pierwszych do wygenerowania: ")
	userN, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(userN))
	if err != nil {
		panic(err)
	}

	lp := 0
	p := 2
	var isFirst bool
	for lp < n {
		isFirst = true
		for i := 2; i <= p-1; i++ {
			if p%i == 0 {
				isFirst = false
				break
			}
		}
		if isFirst {
			fmt.Printf("Liczba pierwsza: %d\n", p)
			lp++
		}
		p++
	}
}
