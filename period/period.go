package period

import (
	"fmt"
	"time"
)

type Period struct {
	from time.Time
	to   time.Time
}

func NewPeriod(from time.Time, to time.Time) *Period {
	return &Period{from: from, to: to}
}

func (p *Period) From() time.Time {
	return p.from
}

func (p *Period) To() time.Time {
	return p.to
}

// Cut - cut period on lesser periods with duration 'len time.Duration'
func Cut(period *Period, len time.Duration) ([]*Period, error) {
	if period.to.Sub(period.from) < len {
		return nil, fmt.Errorf("period должен быть больше или равен len")
	}

	from := period.from
	to := from.Add(len - time.Second)

	var result []*Period
	for {
		if to.After(period.to) {
			to = period.to
			if from.Round(time.Second) == to.Round(time.Second) {
				break
			}

			result = append(result, &Period{from: from, to: to})
			break
		}

		result = append(result, &Period{from: from, to: to})

		from = to.Add(time.Second)
		to = from.Add(len - time.Second)
	}

	return result, nil
}