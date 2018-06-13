package main

import "fmt"

type A struct {
    i int
    j int
}

func (a A) sumFields(k int) int {
    return a.i + a.j + k
}

func main() {
    a := A{i: 2, j: 3}

    fmt.Println(a.sumFields(10))
}