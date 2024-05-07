package main

import (
	"csvreader"
	"fmt"
)

func main() {
	if err := csvreader.Read("FileFrom.csv", "FileTo.txt"); err != nil {
		fmt.Println(err)
	}
}
