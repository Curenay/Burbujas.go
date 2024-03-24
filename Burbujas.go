package main

import (
	"fmt"
	"strconv"
)

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

// Función de ordenación de burbuja
func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	fmt.Println("Introduce hasta 10 números enteros (presiona Enter después de cada número):")
	var input string
	numbers := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Introduce solo números enteros.")
			return
		}
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)
	fmt.Println("Números ordenados:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
