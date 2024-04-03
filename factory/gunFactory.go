package Factory

import (
	"fmt"

	Entity "github.com/PauloJunior5/factory-pattern/guns/entity"
	Interface "github.com/PauloJunior5/factory-pattern/interface"
)

func GetGun(gunType string) (Interface.IGun, error) {

	if gunType == "ak47" {
		return Entity.NewAk47(), nil
	}
	if gunType == "musket" {
		return Entity.NewMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
