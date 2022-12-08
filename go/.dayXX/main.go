package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../data/dayXX")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}