package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	fmt.Print("Liczba rozkładana na czynniki pierwsze: ")

	userP, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	p, err := strconv.Atoi(strings.TrimSpace(userP))
	if err != nil {
		panic(err)
	}

	if p < 1 {
		panic("Liczba nine może być mniejsza od 1")
	}

	var g int = int(math.Sqrt(float64(p)))
	for i := 2; i < g; i++ {
		if p%i == 0 {
			fmt.Printf("Dzielnik: %d\n", i)
			p = p / i
		}

		if p == 1 {
			break
		}
	}

	if p > 1 {
		fmt.Printf("Koniec:%d\n", p)
	}
}
