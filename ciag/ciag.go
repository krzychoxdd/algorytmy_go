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

	fmt.Print("Liczba wyrazów do wygenerowania:")
	userN, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(userN))
	if err != nil {
		panic(err)
	}

	fmt.Print("Pierwszy wyraz ciągu:")
	userA, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.ParseFloat(strings.TrimSpace(userA), 3)
	if err != nil {
		panic(err)
	}

	fmt.Print("Stały przyrost:")
	userD, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	d, err := strconv.ParseFloat(strings.TrimSpace(userD), 3)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= n; i++ {
		fmt.Println(a + (float64(i)-2)*d)
	}
}
