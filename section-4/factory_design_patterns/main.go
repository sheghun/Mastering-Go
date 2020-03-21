package main

import (
	"Mastering-Go/section-4/factory_design_patterns/appliances"
	"fmt"
)

func main() {
	fmt.Println("Entered preferred appliance type")
	fmt.Println("0: Stove ")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var suppliedType int
	fmt.Scan(&suppliedType)

	myAppliance, err := appliances.CreateAppliance(suppliedType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
