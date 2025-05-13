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
		return fmt.Errorf("invalid data format: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("cannot parse steps %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be > 0, got %d", steps)
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	dur, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("cannot parse duration %q: %w", parts[2], err)
	}
	if dur <= 0 {
		return fmt.Errorf("duration must be > 0, got %q", parts[2])
	}
	t.Duration = dur

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Duration <= 0 || t.Weight <= 0 || t.Height <= 0 {
		return "", fmt.Errorf("invalid training data")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var cal float64
	var err error
	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %q", t.TrainingType)
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
		distance,
		speed,
		cal,
	), nil
}
