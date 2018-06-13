package main

import (
    "fmt"
    "strings"
)

// START OMIT
type F func(string) string

func (f F) UpperCase(s string) string {
    return strings.ToUpper(f(s))
}

func (f F) lowerCase(s string) string { // HL
    return strings.ToLower(f(s))
}

func doubleString(a string) string {
    return a + a
}

func main() {
    f := F(doubleString)
    fmt.Println(f("a"))
    fmt.Println(f.UpperCase("a"))
}
// END OMIT