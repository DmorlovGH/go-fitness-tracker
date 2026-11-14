package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные дневной активности и персональные данные пользователя.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m".
// При ошибке данные структуры остаются без изменений.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("данные прогулки должны содержать два поля")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("продолжительность прогулки: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность прогулки должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo возвращает отчёт о дневной прогулке.
// При ошибке расчёта возвращает пустую строку и ошибку.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration,
	)
	if err != nil {
		return "", fmt.Errorf("расчёт дневной активности: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
