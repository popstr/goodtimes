package goodtimes

import (
	"fmt"
	"time"
)

type DateOnly string

func Today() DateOnly {
	n := time.Now()
	d := DateOnlyFromTime(&n)
	return *d
}

func Yesterday() DateOnly {
	n := time.Now().Add(-time.Hour * 24)
	d := DateOnlyFromTime(&n)
	return *d
}

func DateOnlyFromTime(t *time.Time) *DateOnly {
	d := DateOnly(t.Format("2006-01-02"))
	return &d
}

func DateOnlyFromString(s string) (*DateOnly, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}
	d := DateOnly(t.Format("2006-01-02"))
	return &d, nil
}

func (d *DateOnly) String() string {
	if d == nil {
		return ""
	}
	return string(*d)
}

func (d *DateOnly) StringPtr() *string {
	if d == nil {
		return nil
	}
	s := string(*d)
	return &s
}

func (d *DateOnly) Time() time.Time {
	if d == nil {
		return time.Time{}
	}
	t, _ := time.Parse("2006-01-02", string(*d))
	return t
}

func (d *DateOnly) TimePtr() *time.Time {
	if d == nil {
		return nil
	}
	t, _ := time.Parse("2006-01-02", string(*d))
	return &t
}

func (d *DateOnly) Sub(t time.Time) time.Duration {
	return d.Time().Sub(t)
}

func (d *DateOnly) Add(dur time.Duration) *DateOnly {
	t := d.Time().Add(dur)
	return DateOnlyFromTime(&t)
}

func (d *DateOnly) Before(v DateOnly) bool {
	if d == nil {
		return false
	}
	return d.Time().Before(v.Time())
}

func (d *DateOnly) After(v DateOnly) bool {
	if d == nil {
		return false
	}
	return d.Time().After(v.Time())
}

func DatePeriodsOverlapping(s1 *DateOnly, e1 *DateOnly, s2 *DateOnly, e2 *DateOnly) (bool, error) {
	return PeriodsOverlapping(s1.TimePtr(), e1.TimePtr(), s2.TimePtr(), e2.TimePtr())
}
