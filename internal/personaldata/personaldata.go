package personaldata

import "fmt"

// Personal - структура, содержащая персональные данные пользователя.
type Personal struct {
	Name   string  // Имя пользователя.
	Weight float64 // Вес пользователя в килограммах.
	Height float64 // Рост пользователя в сантиметрах.
}

// Print - метод структуры Personal, выводящий данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height) // Вывод имени пользователя.

}
