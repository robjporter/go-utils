package dates

import (
	"regexp"
	"time"
)

type Duration struct {
	time.Duration
}

func MakeDuration(d time.Duration) Duration {
	return Duration{d}
}

const (
	Day   time.Duration = 24 * time.Hour
	Week                = 7 * Day
	Month               = 4 * Week
	Year                = 12 * Month
)

func (d Duration) Days() time.Duration {
	return d.Duration / Day
}

func (d Duration) Weeks() time.Duration {
	return d.Duration / Week
}

func (d Duration) Months() time.Duration {
	return d.Duration / Month
}

func (d Duration) Years() time.Duration {
	return d.Duration / Year
}

func (d Duration) RemainingSecondsStr() string {
	return getNumber((d.Duration % time.Minute / time.Second).String())
}

func (d Duration) RemainingSeconds() time.Duration {
	return d.Duration % time.Minute / time.Second
}

func (d Duration) RemainingMinutesStr() string {
	return getNumber((d.Duration % time.Hour / time.Minute).String())
}

func (d Duration) RemainingMinutes() time.Duration {
	return d.Duration % time.Hour / time.Minute
}

func (d Duration) RemainingHoursStr() string {
	return getNumber((d.Duration % Day / time.Hour).String())
}

func (d Duration) RemainingHours() time.Duration {
	return d.Duration % Day / time.Hour
}

func (d Duration) RemainingDaysStr() string {
	return getNumber((d.Duration % Week / Day).String())
}

func (d Duration) RemainingDays() time.Duration {
	return d.Duration % Week / Day
}

func (d Duration) RemainingWeeksStr() string {
	return getNumber((d.Duration % Month / Week).String())
}

func (d Duration) RemainingWeeks() time.Duration {
	return d.Duration % Month / Week
}

func (d Duration) RemainingMonthsStr() string {
	return getNumber((d.Duration % Year / Month).String())
}

func (d Duration) RemainingMonths() time.Duration {
	return d.Duration % Year / Month
}

func getNumber(s string) string {
	re := regexp.MustCompile("[0-9]+")
	res := re.FindAllString(s, -1)
	if len(res) == 1 {
		return res[0]
	}
	return "0"
}
