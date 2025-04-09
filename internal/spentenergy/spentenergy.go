package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000            // количество метров в километре.
	minInH                     = 60              // количество минут в часе.
	stepLengthCoefficient      = 0.45            // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5             // коэффициент для расчета калорий при ходьбе.
	errInvalidInput            = "invalid input" // сообщение об ошибке
)

// Distance рассчитывает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient  // длина шага
	distanceMeters := float64(steps) * stepLength // дистанция в метрах
	return distanceMeters / mInKm                 // дистанция в километрах
}

// MeanSpeed рассчитывает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	durationHours := duration.Hours()
	return distance / durationHours
}

// RunningSpentCalories рассчитывает потраченные калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New(errInvalidInput)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	return calories, nil
}

// WalkingSpentCalories рассчитывает потраченные калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New(errInvalidInput)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
