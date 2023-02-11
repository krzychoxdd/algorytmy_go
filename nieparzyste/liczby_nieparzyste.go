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

	var i int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pierwsza liczba zakresu: ")

	userA, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.Atoi(strings.TrimSpace(userA))
	i = a
	if err != nil {
		panic(err)
	}

	fmt.Print("Druga liczba zakresu: ")
	userB, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(userB))
	if err != nil {
		panic(err)
	}

	if math.Mod(float64(a), 2) == 0 {
		i++
	}

	for i < b {
		fmt.Println(i)
		i = i + 2
	}
}
