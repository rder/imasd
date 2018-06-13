package main

import (
	"fmt"
)

type Saludator interface{
	saludar(nombre string)
}

type Persona struct{
	nombre string
}

func(this Persona)saludar(nombre string){
	fmt.Println("hola " + nombre + " soy " + this.nombre)
}

func saludarJuanCarlos (saludator Saludator){
	saludator.saludar("Juan Carlos")
}

func main(){
	var jose Persona = Persona{"Jose"}
	jose.saludar("Jorge")
	saludarJuanCarlos(jose)
}