package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training структура для хранения данных о тренировке.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку данных и заполняет поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("неверный формат данных")
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("шаги меньше или равна 0")
	}

	t.TrainingType = data[1]

	if t.Duration, err = time.ParseDuration(data[2]); err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("длительность меньше или равна 0")
	}
	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error // Объявляем переменную для хранения ошибки

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка при расчете калорий для бега: %w", err) // Возвращаем ошибку, если она возникла
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка при расчете калорий для ходьбы: %w", err) // Возвращаем ошибку, если она возникла
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType) // Возвращаем ошибку для неизвестного типа тренировки
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return result, nil // Возвращаем результат и nil (отсутствие ошибки)
}
