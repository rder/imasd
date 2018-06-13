package main

import "fmt"
//START RECOVERY 
func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 1 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
//END RECOVERY

//START BENCHMARK 
func BenchmarkSample(b *testing.B) {
    b.SetBytes(2)
    for i := 0; i < b.N; i++ {
        if x := fmt.Sprintf("%d", 42); x != "42" {
            b.Fatalf("Unexpected string: %s", x)
        }
    }
}
//END BENCHMARK