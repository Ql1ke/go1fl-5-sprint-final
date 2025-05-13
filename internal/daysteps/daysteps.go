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
		return fmt.Errorf("invalid data format: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("cannot parse steps %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be > 0, got %d", steps)
	}
	ds.Steps = steps

	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("cannot parse duration %q: %w", parts[1], err)
	}
	if dur <= 0 {
		return fmt.Errorf("duration must be > 0, got %q", parts[1])
	}
	ds.Duration = dur

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", fmt.Errorf("invalid day steps data")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, cal,
	), nil
}
