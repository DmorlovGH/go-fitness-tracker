package personaldata

import "fmt"

// Personal содержит персональные данные пользователя.
// Вес задаётся в килограммах, рост — в метрах.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит имя, вес и рост пользователя.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n",
		p.Name, p.Weight, p.Height)
}
