package main

import "fmt"
import "time"

func main() {
	showSintaxis()
	show()
	showSwitch()
	showPointers()
}


func showSintaxis() {
	
	var i, j int = 1, 2
	name, power:= "Goku", 1000
	
	fmt.Println(i, j, name, power)

}



func show() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func showSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buen dia!")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes.")
	default:
		fmt.Println("Buenas noches.")
	}
}


func showPointers() {
	
	fmt.Println("##### PUNTEROS #####")
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
