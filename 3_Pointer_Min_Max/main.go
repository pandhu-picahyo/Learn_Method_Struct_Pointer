package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0] // nilai angka index ke 0 masuk ke variabel min
	max = *numbers[0] // nilai angka index ke 0 masuk ke variabel max

	for _, num := range numbers {
		if *num < min {
			min = *num // *num adalah nilai yang baru dimasukkan
		}
		if *num > max {
			max = *num
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Masukkan nilai yang ingin diproses : ")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Printf("%d is the minimum number\n", min)
	fmt.Printf("%d is the maximum number\n", max)
}
