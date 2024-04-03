package Entity

import Interface "github.com/PauloJunior5/factory-pattern/interface"

type musket struct {
	Gun
}

func NewMusket() Interface.IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
