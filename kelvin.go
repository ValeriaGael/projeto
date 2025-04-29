package main

import "fmt"

// declaração da variavel CONST do ponto de ebulição em F
const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK     //variavel do valor da temperatura em graus K
	tempC := (tempK - 273) //Conversão de grau K para ºC //
	fmt.Printf("A temperatura de ebulição da água ºK = %g  e a temperatura de ebulição da água em ºC =  %g", tempK, tempC)

}
