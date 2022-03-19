package main

import (
	"fibonacci-recursive/internal/console"
	"fibonacci-recursive/internal/fibonacci"
	"fmt"
	"time"
)

func main() {

	var userInput int64 = console.GetUserInput()

	recursiveAlgorithmDuration := time.Now()
	recursive := fibonacci.Recursive(userInput, make(map[int64]int64))

	fmt.Println("Recursive -", recursive)
	fmt.Println("Execution time -", time.Since(recursiveAlgorithmDuration))

	fmt.Println("-----------------------")

	iterativeAlgorithmDuration := time.Now()
	iterative := fibonacci.Iterative(userInput)

	fmt.Println("Iterative -", iterative)
	fmt.Println("Execution time -", time.Since(iterativeAlgorithmDuration))

	fmt.Println("-----------------------")

	difference := time.Since(recursiveAlgorithmDuration).Nanoseconds() / time.Since(iterativeAlgorithmDuration).Nanoseconds()

	fmt.Println("Iterative version", difference, "times faster than recursive!")

}
