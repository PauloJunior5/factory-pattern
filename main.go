package main

import (
	"fmt"
	"log"

	Factory "github.com/PauloJunior5/factory-pattern/factory"
)

func main() {
	ak47, err := Factory.GetGun("musket")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ak47)
}
