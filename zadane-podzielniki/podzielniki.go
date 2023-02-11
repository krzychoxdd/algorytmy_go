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
	reader := bufio.NewReader(os.Stdin)

	print("Dzielniki rozdzielone znakiem + np. 15+2009+10: ")
	userMultipliers, err := reader.ReadString('\n')
	println()

	if err != nil {
		panic(err)
	}

	splitedMultipiers := strings.Split(strings.TrimSpace(userMultipliers), "+")
	multipiers := make([]int, len(splitedMultipiers))
	for x, m := range splitedMultipiers {
		multipiers[x], err = strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
	}

	print("Podaj liczbę początątkową zakresu: ")
	userA, err := reader.ReadString('\n')
	println()
	if err != nil {
		panic(err)
	}

	print("Podaj liczbę końcową zakresu: ")
	userB, err := reader.ReadString('\n')
	println()
	if err != nil {
		panic(err)
	}

	println(strings.Repeat("-", 20))

	a, err := strconv.Atoi(strings.TrimSpace(userA))
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(userB))
	if err != nil {
		panic(err)
	}

	var withoutMultipliers []int
	var w int
	var c int
	for i := a; i < b; i++ {
		c = 0
		for _, m := range multipiers {
			if math.Mod(float64(i), float64(m)) == 0 {
				w++
				println(i)
				c = 1
				break
			}
		}
		if c == 0 {
			withoutMultipliers = append(withoutMultipliers, i)
		}
	}

	println(fmt.Sprintf("Znaleziono %d dzielących się liczb", w))

	println("Liczby bez dzielnikow: ")
	g := 1
	for _, y := range withoutMultipliers {
		print(fmt.Sprintf("%d ", y))
		if math.Mod(float64(g), 5) == 0 {
			println()
		}
		g++
	}

	println()
}
