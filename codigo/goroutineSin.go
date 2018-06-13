package main

import "fmt"

func printPares() {
    sum := 0
	for i := 0; i < 800000000; i++ {
		sum += i*2
	}
    fmt.Println(sum)
}

func printImpares() {
    sum := 0
	for i := 0; i < 800000000; i++ {
		sum += i*2+1
	}
    fmt.Println(sum)
}

func main() {
	printPares();
	printPares();
	printPares();
	printPares();
	printPares();
	printImpares();
	printImpares();
	printImpares();
	printImpares();
	printImpares();
}
