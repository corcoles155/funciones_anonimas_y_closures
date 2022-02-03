package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main()  {

	/* Funciones Anónima */

	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	//Podemos redefinir nuestra variable
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2;
	}

	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	//Podemos redefinir nuestra variable
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2;
	}

	fmt.Printf("Dividimos 12 / 3 = %d \n", Calculo(12, 3))

	Operaciones()

	/* Closures */
	tablaDel := 2
	MiTabla := Tabla(tablaDel) //closures

	for i:=1; i<11; i++ {
		//Por cada iteración llamará a la función anónima que hay dentro de la función Tabla(), 
		//guardando en memoria los valores que hay fuera de la función anónima (variables secuencia y numero)
		fmt.Printf(MiTabla())
	}
}	

//Podemos crear funciones anónimas dentro de funciones
func Operaciones()  {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Printf(resultado())
}

//La función que recibe un parámetro y devuelve una función anónima que devuelve un entero.
func Tabla(valor int) func() int  {
	numero := valor
	secuencia := 0
	return func () int {
		//las variables utilizadas secuencia y numero en la función anónima NO son una copia 
		//de las variables de la función Tabla, sino que son las mismas variables.
		secuencia+=1
		return numero * secuencia
	}
}
