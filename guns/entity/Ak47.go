package Entity

import Interface "github.com/PauloJunior5/factory-pattern/interface"

type Ak47 struct {
	Gun
}

func NewAk47() Interface.IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
