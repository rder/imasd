package main
import (
	"os"
)
func main() {
 
	_, err := os.Open("filename.ext")
    if err != nil {
		panic(err) 
		
	}
	
}



