package clock

import "time"

type Clock interface {
	Today(timezone string) (time.Time, error)
}

type systemClock struct {
}

func New() Clock {
	return &systemClock{}
}

func (s systemClock) Today(timezone string) (time.Time, error) {
	if location, err := time.LoadLocation(timezone); err != nil {
		return time.Time{}, err
	} else {
		return time.Now().In(location), nil
	}
}
