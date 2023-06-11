package main

import (
	"fmt"

	"github.com/NikolaM-Dev/go-web-book/chapter2/strcon"
)

func doPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()

	panic("Just panicking for the sake of demo")
	fmt.Println("This will never be called")
}

func main() {
	s := strcon.SwapCase("Gohper")

	fmt.Println("Converted string is :", s)

	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover")
}
