package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	invalidOperator = errors.New("Ingresaste un operador no valido pndjo")
	divideByZero    = errors.New("no se puede dividir por 0, che")
)

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(`ingrese los valores para la calculadora en este orden: \na: primer valor \nb: segundo valor \noperation: puede ser: "suma", "resta", "multiplicacion", "division"`)
	var a, b float64
	var op string
	person1 := Person{}
	fmt.Println("Ingrese el nombre")
	fmt.Scan(&person1.name)
	fmt.Println("Ingrese la edad")
	fmt.Scan(&person1.age)
	fmt.Println("Ingrese los valores para calcular")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&op)

	if strings.Contains(person1.name, "Miguel") || strings.Contains(person1.name, "Jose") {
		fmt.Println(person1.name, " es un nombre feo .l.")
		person1.name = "Alejandro"

	}
	resultado, err := calculadora(float64(a), float64(b), op)

	if err != nil {
		fmt.Println("Algo hiciste mal mmgvo, mira el error: ", err.Error())
		return
	}
	fmt.Println("la persona llamada ", person1.name, " que tiene ", person1.age, " años, calculó el siguiente resultado:", resultado)
}

// Calculadora para brayan
func calculadora(a, b float64, op string) (float64, error) {

	if b == 0 {
		return 0, divideByZero
	}

	switch op {
	case "suma":
		return float64(a + b), nil

	case "resta":
		return float64(a - b), nil

	case "multiplicacion":
		return float64(a * b), nil

	case "division":
		return float64(int(a) / int(b)), nil

	default:
		return 0, invalidOperator
	}
}

var (
	xOutRange = errors.New(" X fuera de rango")
	yOutRange = errors.New(" Y fuera de rango")
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) error {

	if v.X < 2 || v.X > 10 {
		return xOutRange
	}
	if v.Y < 5 && v.Y > 8 {
		return yOutRange
	}
	v.X = v.X * f
	v.Y = v.Y * f
	return nil
}
