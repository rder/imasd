package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	
	dragon := [4]string{"Goku","Vegeta","Beerus","Piccolo"}	
	//Frieza
	var d []string = dragon[:]
	var e []string = dragon[0:2]
	fmt.Println(d)	
	d = append (d, "Frieza")
	//dragon = append (dragon,"Cell")
	fmt.Println(d)
	fmt.Println(dragon)
	e[0] = "Raditz"
	d[4] = "Yamcha"
	d[2] = "Majin Boo"
	
	fmt.Println("Slice d ", d)
	fmt.Println("Slice e ", e)
	fmt.Println("Array Dragon" , dragon)
}