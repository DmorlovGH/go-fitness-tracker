package trainings

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

func TestActionInfoInvalidParameters(t *testing.T) {
	tests := []struct {
		name     string
		steps    int
		weight   float64
		height   float64
		duration time.Duration
	}{
		{"нулевые шаги", 0, 75, 1.75, time.Hour},
		{"отрицательные шаги", -1, 75, 1.75, time.Hour},
		{"нулевой вес", 1000, 0, 1.75, time.Hour},
		{"отрицательный вес", 1000, -75, 1.75, time.Hour},
		{"нулевой рост", 1000, 75, 0, time.Hour},
		{"отрицательный рост", 1000, 75, -1.75, time.Hour},
		{"нулевая продолжительность", 1000, 75, 1.75, 0},
		{"отрицательная продолжительность", 1000, 75, 1.75, -time.Hour},
	}

	for _, trainingType := range []string{"Бег", "Ходьба"} {
		t.Run(trainingType, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					training := Training{
						Steps:        tt.steps,
						TrainingType: trainingType,
						Duration:     tt.duration,
						Personal: personaldata.Personal{
							Weight: tt.weight,
							Height: tt.height,
						},
					}

					info, err := training.ActionInfo()
					if err == nil {
						t.Error("ожидалась ошибка расчёта")
					}
					if info != "" {
						t.Errorf("ожидалась пустая строка, получено %q", info)
					}
				})
			}
		})
	}
}
