package main

import "fmt"

type Car struct {
	Type string
	Fuel float64
}

func (car *Car) PerkiraanJarakTempuh() {
	var efisiensi float64
	switch car.Type {
		case "pertamax":
			efisiensi = 2
		case "pertalite":
			efisiensi = 1.5
		case "solar":
			efisiensi = 1.2
	}	
	fmt.Println(car.Fuel * efisiensi, "mill")
}

func main() {
	var pertaliteCar = Car{"pertalite", 50}
	var pertamaxCar = Car{"pertamax", 100}
	
	pertaliteCar.PerkiraanJarakTempuh()
	pertamaxCar.PerkiraanJarakTempuh()
}