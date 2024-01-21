package main

import (
	"errors"
	"fmt"
	"math"
)

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
	if v.Y < 5 || v.Y > 8 {
		return yOutRange
	}
	v.X = v.X * f
	v.Y = v.Y * f
	return nil
}

func main() {

	v := Vertex{2, 20}
	err := v.Scale(10)
	if err != nil {
		fmt.Println("Algo hiciste mal mmgvo, mira el error: ", err.Error())
		return
	}
	fmt.Println(v.Abs())

}
