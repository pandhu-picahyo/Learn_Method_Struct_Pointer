package main

import "fmt"

// struct Car dengan property Type dan FuelIn
type Car struct {
	Type   string  // string karena huruf
	FuelIn float32 // float karena angka desimal
}

// Method untuk menghitung jarak berdasarkan bahan bakar yang terisi
func (car Car) PerkiraanJarak() float32 {
	bensinPerMill := 1.5
	jarak := car.FuelIn / float32(bensinPerMill)
	return jarak
}

func main() {
	var car Car
	car.Type = "BMW"
	car.FuelIn = 60

	// %.2f adalah format supaya 2 angka dibelakang koma
	fmt.Printf("Perkiraan jarak yang dapat ditempuh adalah : %.2f Mill", car.PerkiraanJarak())
}
