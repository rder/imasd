package main

import "time"
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
	go printPares();
	go printPares();
	go printPares();
	go printPares();
	go printPares();
	go printImpares();
	go printImpares();
	go printImpares();
	go printImpares();
    go printImpares();
    time.Sleep(10000 * time.Millisecond)
}
