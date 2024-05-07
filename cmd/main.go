package main

import (
	"csvreader"
	"fmt"
)

func main() {
	if err := csvreader.Read("file.csv", "FileTo.txt"); err != nil {
		fmt.Println(err)
	}
}
