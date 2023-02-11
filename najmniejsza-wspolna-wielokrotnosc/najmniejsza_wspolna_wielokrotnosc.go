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

	fmt.Print("A: ")

	userA, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	a, err := strconv.Atoi(strings.TrimSpace(userA))
	if err != nil {
		panic(err)
	}

	fmt.Print("A: ")

	userB, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(strings.TrimSpace(userB))
	if err != nil {
		panic(err)
	}

	axb := a * b
	for a != b {
		if a < b {
			b = b - a
		} else {
			a = a - b
		}
	}
	nwd := a

	result := axb / nwd

	fmt.Println(fmt.Sprintf("NWW: %d", result))
}
