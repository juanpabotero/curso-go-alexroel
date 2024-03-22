package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"rsc.io/quote"

	"github.com/juanpabotero/greeting" // importar un paquete propio
)

// declaracion a nivel de paquete
const x = 0xFF // hexadecimal -> 255

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())

	// variables
	var age uint8 = 255
	name := "Juan"
	fmt.Println(name, age)

	// constantes
	const pi = 3.1416
	fmt.Println(x)

	// se puede usar iota para crear constantes con valores exponenciales
	const (
		_  = iota
		KB = 1 << (iota * 10)
		MB = 1 << (iota * 10)
		GB = 1 << (iota * 10)
	)
	fmt.Println(KB, MB, GB)

	// byte
	var a byte = 'a' // 97 (ASCII)
	s := "Hello"
	b := s[0] // 72 (ASCII)
	// rune
	var r rune = '♥' // 9829 (Unicode)

	fmt.Println(a, b, r)

	var (
		boolean  bool
		num      int
		myString string
	)
	fmt.Println(boolean, num, myString)

	// conversion de tipos
	temperatureString := "88"
	// el segndo valor que devuelve es un error, si no se logra hacer la conversion
	temperatureInt, _ := strconv.Atoi(temperatureString)
	fmt.Println("temperatura", temperatureInt)

	// fmt
	// var nameScan string
	// fmt.Println("Ingrese su nombre:")
	// se pone la referencia de la variable que va a recibir el valor
	// fmt.Scanln(&nameScan)
	// fmt.Println("Hola", nameScan)

	// math
	var a1, b1 int = 5, 3
	fmt.Println(a1 + b1)
	fmt.Println("IsNaN?", math.IsNaN(5.0))

	// time
	t := time.Now()
	fmt.Println(t)
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().Add(time.Hour * 24))
	// fmt.Println(time.Now().Sub(time.Now().Add(time.Hour * 24)))

	// ciclo
	for i := 0; i < 5; i++ {
		if i == 4 {
			break
		}
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	arr := []int{1, 2, 3}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// array
	arr2 := [...]int{1, 2, 3}
	fmt.Println("array", arr2)
	arr3 := [3]int{1}
	fmt.Println("array", arr3)
	// matriz
	matriz := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("matriz", matriz)

	// slice
	arr4 := [3]int{1, 2, 3}
	// se puede crear un slice a partir de un array
	slice := arr4[:]
	// se puede crear un slice a partir de un array con un rango
	slice2 := arr4[1:3] // [2, 3]
	slice3 := append(slice2, 4, 5, 6)
	slice4 := append(slice3[:2], slice3[3:]...)
	fmt.Println("slice", slice, slice2, slice3, slice4)

	// poniter
	x := 10
	y := &x
	fmt.Println("pointer", y, *y)

	// random
	rand.Intn(100) // numero aleatorio entre 0 y 100
	rand.Float64() // numero aleatorio entre 0 y 1
	fmt.Println(rand.Intn(100))

	// usar modulo propio
	message, err := greeting.Hello("Juan")
	fmt.Println("greeting", message, err)

	names := []string{"Juan", "Pedro", "Maria"}
	messages, err := greeting.Hellos(names)
	fmt.Println("messages", messages, err)

	// funcion variadica
	PrintList(1, "Juan", true, 4, "Pablo") // 1, Juan, true, 4, Pablo

	fmt.Println(Sum[int](1, 2, 3, 4, 5))               // 15
	fmt.Println(Sum[float64](1.5, 2.5, 3.5, 4.5, 5.5)) // 17.5

	// excercise 1
	// calculateTriangle()
}

// funcion variadica
func PrintList(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func Sum[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// excercise 1
func calculateTriangle() {
	var width, height float64
	const numberDecimals = 2

	fmt.Println("Ingrese el ancho:")
	fmt.Scanln(&width)
	fmt.Println("Ingrese el alto:")
	fmt.Scanln(&height)

	area := (width * height) / 2
	hipo := math.Sqrt(math.Pow(width, 2) + math.Pow(height, 2))
	perimetro := width + height + hipo

	fmt.Printf("Area: %.*f \n", numberDecimals, area)
	fmt.Printf("Perimetro: %.*f", numberDecimals, perimetro)
}
