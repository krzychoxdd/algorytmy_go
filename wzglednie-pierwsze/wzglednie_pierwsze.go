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

	fmt.Print("Pierwsza liczba zakresu([a, b]) :")
	userA, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.Atoi(strings.TrimSpace(userA))
	if err != nil {
		panic(err)
	}

	fmt.Print("Druga liczba zakresu([a, b]) :")
	userB, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(userB))
	if err != nil {
		panic(err)
	}

	fmt.Print("Podzielnik :")
	userM, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	m, err := strconv.Atoi(strings.TrimSpace(userM))
	if err != nil {
		panic(err)
	}

	for i := a; i < b; i++ {
		ax := i
		bx := m
		for ax != bx {
			if ax < bx {
				bx = bx - ax
			} else {
				ax = ax - bx
			}
		}

		if ax == 1 {
			fmt.Println(i)
		}
	}
}
