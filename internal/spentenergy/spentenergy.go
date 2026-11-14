package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент длины шага.
	walkingCaloriesCoefficient = 0.5  // коэффициент калорий при ходьбе.
)

// WalkingSpentCalories возвращает количество калорий,
// потраченных при ходьбе. Все входные параметры должны быть положительными.
func WalkingSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, fmt.Errorf("расчёт калорий при ходьбе: %w", err)
	}

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий,
// потраченных при беге. Все входные параметры должны быть положительными.
func RunningSpentCalories(
	steps int,
	weight, height float64,
	duration time.Duration,
) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	speed := MeanSpeed(steps, height, duration)
	calories := weight * speed * duration.Minutes() / minInH

	return calories, nil
}

// MeanSpeed возвращает среднюю скорость в километрах в час.
// При неположительных шагах, росте или продолжительности возвращает ноль.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает пройденную дистанцию в километрах.
// Рост пользователя задаётся в метрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
