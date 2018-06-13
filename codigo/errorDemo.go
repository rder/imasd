package main
import (
	"fmt"
	"os"
)
func main() {
	i, err := os.Open("filename.ext")
	if err != nil {
		fmt.Printf("No existe el archivo: %v\n", err);
		return
	}
	fmt.Println("salio todo ok:", i);
}
