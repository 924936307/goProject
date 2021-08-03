package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("hello,world!")
}

func digits(number int, dchan chan int) {
	for number != 0 {
		param := number % 10
		dchan <- param
		number = number / 10
	}
	close(dchan)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for c := range dch {
		sum += c * c
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	tempChan := make(chan int)
	go digits(number, tempChan)
	for i := range tempChan {
		sum += i * i * i
	}
	cubeop <- sum
}

func TestCalc(m *testing.T) {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Printf("results:param1:%d,param2:%d \n", squares, cubes)
	fmt.Println("Final output", squares+cubes)
}
