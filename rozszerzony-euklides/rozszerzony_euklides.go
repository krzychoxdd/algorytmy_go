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

	fmt.Print("Liczba naturalna której odwrotność modułu obliczamy:")
	userA, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.Atoi(strings.TrimSpace(userA))
	if err != nil {
		panic(err)
	}

	fmt.Print("Moduł odwrotności:")
	userB, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(userB))
	if err != nil {
		panic(err)
	}

	var u, w, x, z, q int
	u = 1
	w = a
	x = 0
	z = b

	for w != 0 {
		if w < z {
			q = u
			u = x
			x = q
			q = w
			w = z
			z = q
		}
		q = w / z
		u = u - (q * x)
		w = w - (q * z)
	}

	if z == 1 {
		if x < 0 {
			x = x + b
		}
		fmt.Printf("Wynik: %d", x)
	} else {
		fmt.Print("Brak wyniku")
	}
}
