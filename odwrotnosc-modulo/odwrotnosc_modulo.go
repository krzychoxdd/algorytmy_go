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

	fmt.Print("Podaj wartość A: ")
	userA, _ := rd.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(userA))

	fmt.Print("Podaj wartość B: ")
	userB, _ := rd.ReadString('\n')
	b, _ := strconv.Atoi(strings.TrimSpace(userB))

	var x int
	for i := 1; x != 1; i++ {
		x = a * i % b
		if x == 1 {
			fmt.Println(i)
		}
	}
}
