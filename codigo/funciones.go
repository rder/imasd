package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func valores () {
  xs := []float64{98,93,77,82,83}
  total := 0.0
  for _, v := range xs {
    total += v
  }
  fmt.Println(total / float64(len(xs)))
}

func retornoMultiple() (int, int) {
  return 5, 6
}

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
	v := Vertex{3, 4}
	x, y := retornoMultiple()
	Scale(&v, 10)
	fmt.Println(Abs(v))
	fmt.Println(" x,y " )
}