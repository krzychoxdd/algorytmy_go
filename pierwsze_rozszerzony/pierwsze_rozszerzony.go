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

	fmt.Print("Ilość liczb pierwszych do wygenerowania: ")
	userN, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(userN))

	if err != nil {
		panic(err)
	}

	var lp, k, d, p int = 0, 1, 1, 2
	var tlp []int

	for lp < n {
		check := true
		if lp < 3 {
			p = p + lp
		} else {
			p = 6*k + d
			if d == 1 {
				d = -1
				k = k + 1
			} else {
				d = 1
			}

			g := int(math.Sqrt(float64(p)))
			for i := 2; tlp[i] <= g; i++ {
				if p%tlp[i] == 0 {
					check = false
					break
				}
			}
		}

		if check {
			tlp = append(tlp, p)
			lp = lp + 1
			fmt.Printf("Liczba pierwsza: %d\n", p)
		}
	}
}
