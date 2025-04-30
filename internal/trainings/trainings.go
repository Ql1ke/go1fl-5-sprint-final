package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат данных: %q", datastring)
	}

	s := parts[0]
	if strings.TrimSpace(s) != s || s == "" {
		return fmt.Errorf("неверное количество шагов: %q", s)
	}
	steps, err := strconv.Atoi(s)
	if err != nil || steps <= 0 {
		return fmt.Errorf("неверное количество шагов: %q", s)
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	dur := parts[2]
	d, err := time.ParseDuration(dur)
	if err != nil || d <= 0 {
		return fmt.Errorf("неверная продолжительность: %q", dur)
	}
	t.Duration = d

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 ||
		t.Duration <= 0 ||
		t.Weight <= 0 ||
		t.Height <= 0 {
		return "", fmt.Errorf("некорректные данные для тренировки")
	}

	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var cal float64
	var err error
	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %q", t.TrainingType)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		cal,
	), nil
}
