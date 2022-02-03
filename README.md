# Funciones Anónimas y Closures en GO

## Funciones Anónima
Las funciones anónimas son funciones que no tienen nombre, con las que se pueden formar closures. Podemos modificarlas en tiempo de ejecución.

```var Calculo func```

```
func Tabla(valor int) func() int  {
	numero := valor
	secuencia := 0
	return func () int {
		secuencia+=1
		return numero * secuencia
	}
}
```

## Closures
Es una implementación de una función anónima como retorno de una función principal, permite que la función principal pueda ser asignada o delegada a una variable que es en sí el closure. En ese momento, las variables de la función principal se registran o inicializan y el control pasa al closure.
Cuando se invoca al closure, internamente llama a la sentencia return, que es la que tiene la función anónima con el bloque de código de la operativa.
```
//closure
MiTabla := Tabla(2)

for i:=1; i<11; i++ {
	fmt.Printf(MiTabla())
}

//Función que retorna una funcion anonima
func Tabla(valor int) func() int  {
	numero := valor
	secuencia := 0
	return func () int {
		secuencia+=1
		return numero * secuencia
	}
}
```
También podemos utilizar parámetros en nuestros closures. Para ello, los parámetros se definen en la función de retorno.
```
//closure
acumulador := acumular(2)

fmt.Printf(acumulador(4)) //El resultado es 4
fmt.Printf(acumulador(6)) //El resultado es 10 (4+6)
fmt.Printf(acumulador(20)) //El resultado es 30 (10+20)

//Función que retorna una funcion anonima a la que le pasamos un parámetro
func acumular() func(incremento int) int  {
	resultado := 0
	return func (incremento int) int {
		return resultado+=incremento
	}
}
```
