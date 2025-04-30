package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
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
	ds.Steps = steps

	dur := parts[1]
	d, err := time.ParseDuration(dur)
	if err != nil || d <= 0 {
		return fmt.Errorf("неверная продолжительность: %q", dur)
	}
	ds.Duration = d

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 ||
		ds.Duration <= 0 ||
		ds.Weight <= 0 ||
		ds.Height <= 0 {
		return "", fmt.Errorf("некорректные данные для прогулки")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, dist, cal,
	), nil
}
