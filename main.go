package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	var max int

	flag.IntVar(&max, "max", 100, "random numbers between 0 and this")
	flag.Parse()
	totalCorrect := 0
	totalWrong := 0

	for {
		answer := -1
		num1 := rand.Intn(max) + 1
		num2 := rand.Intn(max) + 1

		fmt.Println()
		fmt.Println()
		fmt.Println()
		for {
			fmt.Printf("%v + %v = ", num1, num2)

			sanswer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Print(err)
			}
			answer, err = strconv.Atoi(strings.TrimSpace(sanswer))
			if err != nil {
				fmt.Print(err)
			}

			if num1+num2 == answer {
				fmt.Println("Correct!")
				totalCorrect = totalCorrect + 1
				fmt.Printf("You have answered %v correctly.\n", totalCorrect)
				fmt.Printf("You have answered %v incorrectly.\n", totalWrong)
				break
			} else {
				fmt.Println("Sorry, Try Again!")
				totalWrong = totalWrong + 1
			}
		}
	}
}
