package times

import (
	"time"

	"strings"

	"github.com/robjporter/go-functions/as"
)

const (
	VERSION = "1.2.0.0"
)

type Times struct {
	year     int
	month    int
	day      int
	hour     int
	minute   int
	second   int
	timezone *time.Location
	err      error
	format   string
	reset    bool
}

type Diff struct {
	year  int
	month int
	week  int
	day   int
	hour  int
	min   int
	sec   int
}

func New(year, month, day, hour, minute, second int, location string) *Times {
	loc, err := time.LoadLocation(location)
	if err == nil {

		t := Times{
			year:     year,
			month:    month,
			day:      day,
			hour:     hour,
			minute:   minute,
			second:   second,
			timezone: loc,
			err:      nil,
			format:   "2006-01-02 15:04:05",
			reset:    true,
		}
		return &t
	} else {
		return nil
	}
}

func NewToday(location string) *Times {
	ti := time.Now()
	t := Times{
		year:     ti.Year(),
		month:    MonthToNumber(ti.Month()),
		day:      ti.Day(),
		hour:     ti.Hour(),
		minute:   ti.Minute(),
		second:   ti.Second(),
		timezone: ti.Location(),
		err:      nil,
		format:   "2006-01-02 15:04:05",
		reset:    true,
	}
	return &t

}

func NewTodayAuto() *Times {
	ti := time.Now()
	t := Times{
		year:     ti.Year(),
		month:    MonthToNumber(ti.Month()),
		day:      ti.Day(),
		hour:     ti.Hour(),
		minute:   ti.Minute(),
		second:   ti.Second(),
		timezone: time.Now().Location(),
		err:      nil,
		format:   "2006-01-02 15:04:05",
		reset:    true,
	}
	return &t

}

///////////////// TO IMPLEMENT /////////////////
func (t *Times) WeekNumber() int {
	return 0
}

func (t *Times) GetStartOfWeekNumber(week int) time.Weekday {
	return time.Monday
}

func (t *Times) GetDateForStartOfWeekNumber(week int) time.Time {
	return time.Now()
}

func (t *Times) Location() string {
	return t.timezone.String()
}

func (t *Times) TimeToNextQuarter() (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) FirstInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) SecondInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) ThirdInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) FourthInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) FifthInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) LastInMonth(day time.Weekday) (time.Time, error) {
	return time.Now(), nil
}

func (t *Times) NextEclipse() (time.Time, error) {
	// https://www.timeanddate.com/eclipse/list.html
	return time.Now(), nil
}

func (t *Times) TimeZoneDiff(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffYears(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffMonths(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffWeeks(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffDays(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffHours(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffMinutes(tz time.Location) string {
	return ""
}
func (t *Times) TimeZoneDiffSeconds(tz time.Location) string {
	return ""
}

///////////////// HELPERS /////////////////
func (t *Times) GetSecond() int {
	return t.second
}

func (t *Times) GetMinute() int {
	return t.minute
}

func (t *Times) GetHour() int {
	return t.hour
}

func (t *Times) GetDay() int {
	return t.day
}

func (t *Times) GetMonth() int {
	return t.month
}

func (t *Times) GetYear() int {
	return t.year
}

func (t *Times) GetMonthName() string {
	return MonthIntToName(t.GetMonth())
}

func processNumber(num string) string {
	if len(num) == 1 {
		return "0" + num
	}
	return num
}

func (t *Times) buildTimeString() string {
	dat := as.ToString(t.year) + "-" + processNumber(as.ToString(t.month)) + "-" + processNumber(as.ToString(t.day))
	tim := as.ToString(t.hour) + ":" + processNumber(as.ToString(t.minute)) + ":" + processNumber(as.ToString(t.second))
	return dat + " " + tim
}

func (t *Times) formattedDate() (time.Time, error) {
	return time.Parse(t.format, t.buildTimeString())
}

func MonthToNumber(month time.Month) int {
	switch month {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 0
}

func MonthIntToName(month int) string {
	monthNames := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	return monthNames[month-1]
}

func MonthToName(month time.Month) string {
	monthNames := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	return monthNames[MonthToNumber(month)-1]
}

func MonthNameToNumber(month string) int {
	monthLong := []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	monthShort := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	for i := 0; i < len(monthLong); i++ {
		if doesMatchMonth(month, monthShort[i], monthLong[i]) {
			return i + 1
		}
	}
	return -1
}

func MonthNameToFullName(month string) string {
	monthNames := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	monthLong := []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	monthShort := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	for i := 0; i < len(monthLong); i++ {
		if doesMatchMonth(month, monthShort[i], monthLong[i]) {
			return monthNames[i]
		}
	}
	return ""
}

func doesMatchMonth(month, monthShort, monthLong string) bool {
	month = strings.ToLower(month)
	if month == monthLong || month == monthShort {
		return true
	} else {
		if strings.HasPrefix(month, monthShort) {
			return true
		}
	}
	return false
}

func (t *Times) getWeekDayNumber(day time.Weekday) int {
	switch day {
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	case time.Sunday:
		return 7
	}
	return 0
}

func (t *Times) maxDaysInMonth() int {
	switch t.month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if t.IsLeapYear() {
			return 29
		} else {
			return 28
		}
	}
	return 0
}

func (t *Times) updateValues(date time.Time) {
	t.year = date.Year()
	t.month = MonthToNumber(date.Month())
	t.day = date.Day()
	t.hour = date.Hour()
	t.minute = date.Minute()
	t.second = date.Second()
}

func (t *Times) formatDifference(data time.Duration) Diff {
	tmp := Diff{}
	splits := strings.Split(as.ToString(data), "h")
	tmp.hour = 0
	if len(splits) > 1 {
		tmp.hour = as.ToInt(splits[0])
	}
	splits = strings.Split(as.ToString(data), "m")
	splits2 := strings.Split(splits[0], "h")
	tmp.min = 0
	if len(splits2) > 1 {
		tmp.min = as.ToInt(splits2[1])
	}
	tmp.sec = 0
	splits = strings.Split(as.ToString(data), "s")
	splits2 = strings.Split(splits[0], "m")
	if len(splits2) > 1 {
		tmp.sec = as.ToInt(splits2[1])
	}

	tmp.day = 0
	if tmp.hour > 23 {
		tmp.day = tmp.hour / 24
		tmp.hour = tmp.hour % 24
	}

	tmp.week = 0
	if tmp.day > 6 {
		tmp.week = tmp.day / 7
		tmp.day = tmp.day % 7
	}

	tmp.month = 0
	if tmp.week > 4 {
		tmp.month = tmp.week / 4
		tmp.week = tmp.week % 4
	}

	tmp.year = 0
	if tmp.month > 11 {
		tmp.year = tmp.month / 12
		tmp.month = tmp.month % 12
	}

	return tmp
}

func (t *Times) processStruct(tmp Diff) string {
	result := ""
	if tmp.sec > 0 {
		result = as.ToString(tmp.sec) + "s"
	}
	if tmp.min > 0 {
		result = as.ToString(tmp.min) + "m" + result
	}
	if tmp.hour > 0 {
		result = as.ToString(tmp.hour) + "h" + result
	}
	if tmp.day > 0 {
		result = as.ToString(tmp.day) + "d" + result
	}
	if tmp.week > 0 {
		result = as.ToString(tmp.week) + "w" + result
	}
	if tmp.month > 0 {
		result = as.ToString(tmp.month) + "M" + result
	}
	if tmp.year > 0 {
		result = as.ToString(tmp.year) + "y" + result
	}
	return result
}

///////////////// IMPLEMENTED /////////////////
func (t *Times) TaxYear() string {
	if t.month < 4 && t.day < 6 {
		return as.ToString(t.year-1) + "-" + as.ToString(t.year)
	} else if t.month > 3 && t.day > 5 {
		return as.ToString(t.year) + "-" + as.ToString(t.year+1)
	}
	return ""
}

func (t *Times) StartOfTaxYear() (time.Time, error) {
	t2 := *t
	if t.month < 4 && t.day < 6 {
		t2.year -= 1
	}

	t2.month = 4
	t2.day = 6
	t2.hour = 00
	t2.minute = 00
	t2.second = 00
	return t2.formattedDate()
}

func (t *Times) EndOfTaxYear() (time.Time, error) {
	t2 := *t
	if t.month > 3 && t.day > 5 {
		t2.year += 1
	}

	t2.month = 4
	t2.day = 5
	t2.hour = 23
	t2.minute = 59
	t2.second = 59
	return t2.formattedDate()
}

func (t *Times) TimeToTaxYear() string {
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 3 && t.day > 5 {
		year += 1
	}
	t2 := New(year, 4, 6, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) TimeToTaxYearDiff() Diff {
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 3 && t.day > 5 {
		year += 1
	}
	t2 := New(year, 4, 6, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) AddDecade() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(10, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddDecades(decades int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(decades*10, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDecade() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-10, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDecades(decades int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-decades*10, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) StartOfQuarter() (time.Time, error) {
	t2 := *t

	if t.Is1stQuarter() {
		t2.month = 1
		t2.day = 1
	} else if t.Is2ndQuarter() {
		t2.month = 4
		t2.day = 1
	} else if t.Is3rdQuarter() {
		t2.month = 7
		t2.day = 1
	} else if t.Is4thQuarter() {
		t2.month = 10
		t2.day = 1
	}

	t2.hour = 00
	t2.minute = 00
	t2.second = 00

	return t2.formattedDate()
}

func (t *Times) EndOfQuarter() (time.Time, error) {
	t2 := *t

	if t.Is1stQuarter() {
		t2.month = 3
		t2.day = 31
	} else if t.Is2ndQuarter() {
		t2.month = 6
		t2.day = 30
	} else if t.Is3rdQuarter() {
		t2.month = 9
		t2.day = 31
	} else if t.Is4thQuarter() {
		t2.month = 12
		t2.day = 31
	}

	t2.hour = 23
	t2.minute = 59
	t2.second = 59

	return t2.formattedDate()
}

func (t *Times) Quarter() string {
	if t.Is1stQuarter() {
		return "1st"
	} else if t.Is2ndQuarter() {
		return "2nd"
	} else if t.Is3rdQuarter() {
		return "3rd"
	} else if t.Is4thQuarter() {
		return "4th"
	}
	return ""
}

func (t *Times) QuarterNumber() int {
	if t.Is1stQuarter() {
		return 1
	} else if t.Is2ndQuarter() {
		return 2
	} else if t.Is3rdQuarter() {
		return 3
	} else if t.Is4thQuarter() {
		return 4
	}
	return 0
}

func (t *Times) Is1stQuarter() bool {
	start := New(t.year, 1, 1, 00, 00, 00, t.timezone.String())
	end := New(t.year, 3, 31, 23, 59, 59, t.timezone.String())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) Is2ndQuarter() bool {
	start := New(t.year, 4, 1, 00, 00, 00, t.timezone.String())
	end := New(t.year, 6, 30, 23, 59, 59, t.timezone.String())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) Is3rdQuarter() bool {
	start := New(t.year, 7, 1, 00, 00, 00, t.timezone.String())
	end := New(t.year, 9, 31, 23, 59, 59, t.timezone.String())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) Is4thQuarter() bool {
	start := New(t.year, 10, 1, 00, 00, 00, t.timezone.String())
	end := New(t.year, 12, 31, 23, 59, 59, t.timezone.String())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) NextLeapYear() string {
	year := t.year
	for i := 0; i < 6; i++ {
		if internalIsLeapYear(year) {
			return as.ToString(year)
		} else {
			year += 1
		}
	}

	return ""
}

func (t *Times) TimeToSpring() string {
	// 20th March
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 3 {
		year += 1
	}
	t2 := New(year, 3, 20, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) TimeToSpringDiff() Diff {
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 3 {
		year += 1
	}
	t2 := New(year, 3, 20, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) TimeToSummer() string {
	// 21st June
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 6 {
		year += 1
	}
	t2 := New(year, 6, 21, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) TimeToSummerDiff() Diff {
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 6 {
		year += 1
	}
	t2 := New(year, 6, 21, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) TimeToAutumn() string {
	// 22nd September
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 6 {
		year += 1
	}
	t2 := New(year, 9, 22, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) TimeToAutumnDiff() Diff {
	tmp, _ := t.formattedDate()
	year := t.year
	if t.month > 6 {
		year += 1
	}
	t2 := New(year, 9, 22, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) TimeToWinter() string {
	// 21st December
	// Year before leap year = 22nd
	tmp, _ := t.formattedDate()
	nextLeap := as.ToInt(t.NextLeapYear())

	day := 21
	year := t.year
	if t.month > 6 {
		year += 1
	}

	if nextLeap-year == 1 {
		day = 22
	}

	t2 := New(year, 12, day, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) TimeToWinterDiff() Diff {
	tmp, _ := t.formattedDate()
	nextLeap := as.ToInt(t.NextLeapYear())

	day := 21
	year := t.year
	if t.month > 6 {
		year += 1
	}

	if nextLeap-year == 1 {
		day = 22
	}

	t2 := New(year, 12, day, 00, 00, 00, t.Location())
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) IsSpring() bool {
	// 20th March
	start := New(t.year, 3, 20, 00, 00, 00, t.Location())
	end := New(t.year, 6, 20, 23, 59, 59, t.Location())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) IsSummer() bool {
	// 21st June
	start := New(t.year, 6, 21, 00, 00, 00, t.Location())
	end := New(t.year, 9, 21, 23, 59, 59, t.Location())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) IsAutumn() bool {
	// 22nd September
	start := New(t.year, 9, 22, 00, 00, 00, t.Location())
	end := New(t.year, 9, 21, 23, 59, 59, t.Location())

	if t.IsBetween(start, end) {
		return true
	}
	return false
}

func (t *Times) IsWinter() bool {
	// 21st December
	// Year before leap year = 22nd
	start1 := New(t.year, 1, 1, 00, 00, 00, t.Location())
	end1 := New(t.year, 3, 19, 23, 59, 59, t.Location())
	start2 := New(t.year, 12, 21, 00, 00, 00, t.Location())
	end2 := New(t.year, 12, 31, 23, 59, 59, t.Location())

	if t.IsBetween(start1, end1) || t.IsBetween(start2, end2) {
		return true
	}
	return false
}

func (t *Times) Season() string {
	if t.IsSpring() {
		return "Spring"
	} else if t.IsSummer() {
		return "Summer"
	} else if t.IsAutumn() {
		return "Autumn"
	} else {
		return "Winter"
	}
}

func (t *Times) Copy() *Times {
	t2 := *t
	return &t2
}

func (t *Times) Difference(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	tmp4 := t.formatDifference(tmp3)
	return t.processStruct(tmp4)
}

func (t *Times) DifferenceDiff(t2 *Times) Diff {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)

	return t.formatDifference(tmp3)
}

func (t *Times) DiffInSeconds(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Seconds())
}

func (t *Times) DiffInMinutes(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Minutes())
}

func (t *Times) DiffInHours(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Hours())
}

func (t *Times) DiffInDays(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(tmp3.Hours() / 24)
}

func (t *Times) DiffInWeeks(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString((tmp3.Hours() / 24) / 7)
}

func (t *Times) DiffInMonths(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString(((tmp3.Hours() / 24) / 7) / 4)
}

func (t *Times) DiffInYears(t2 *Times) string {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	tmp3 := tmp2.Sub(tmp)
	return as.ToString((((tmp3.Hours() / 24) / 7) / 4) / 12)
}

func (t *Times) IsFuture(t2 *Times) bool {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	if tmp2.After(tmp) {
		return true
	}

	return false
}

func (t *Times) IsBetween(t2 *Times, t3 *Times) bool {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()
	tmp3, _ := t3.formattedDate()

	if tmp.After(tmp2) && tmp.Before(tmp3) {
		return true
	}
	return false
}

func (t *Times) IsPast(t2 *Times) bool {
	tmp, _ := t.formattedDate()
	tmp2, _ := t2.formattedDate()

	if tmp.Before(tmp2) {
		return true
	}

	return false
}

func (t *Times) ISOWeek() (string, error) {
	tmp, _ := t.formattedDate()
	_, week := tmp.ISOWeek()
	return as.ToString(week), nil
}

func (t *Times) Format1() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format2() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format3() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/Jan/2006 03:04:05 PM")
	return result, err
}

func (t *Times) Format4() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/Jan/2006 15:04:05")
	return result, err
}

func (t *Times) Format5() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Mon Jan")
	return result, err
}

func (t *Times) Format6() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Monay January")
	return result, err
}

func (t *Times) Format7() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM Jan")
	return result, err
}

func (t *Times) Format8() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("2/1/6 3:4:5 PM")
	return result, err
}

func (t *Times) Format9() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("_2/1/6 3:4:5 PM")
	return result, err
}

func (t *Times) Format10() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/06 03:04:05 PM")
	return result, err
}

func (t *Times) Format11() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05 PM")
	return result, err
}

func (t *Times) Format12() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000 PM")
	return result, err
}

func (t *Times) Format13() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000000 PM")
	return result, err
}

func (t *Times) Format14() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 03:04:05.000000000 PM")
	return result, err
}

func (t *Times) Format15() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 MST")
	return result, err
}

func (t *Times) Format16() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z7")
	return result, err
}

func (t *Times) Format17() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z07")
	return result, err
}

func (t *Times) Format18() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z0700")
	return result, err
}

func (t *Times) Format19() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 Z07:00")
	return result, err
}

func (t *Times) Format20() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format("02/01/2006 15:04:05 -07:00")
	return result, err
}

func (t *Times) Format822() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC822)
	return result, err
}

func (t *Times) Format1123() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC1123)
	return result, err
}

func (t *Times) Format1123z() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC1123Z)
	return result, err
}

func (t *Times) Format3339() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC3339)
	return result, err
}

func (t *Times) Format3339nano() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC3339Nano)
	return result, err
}

func (t *Times) Format8222z() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC822Z)
	return result, err
}

func (t *Times) Format850() (string, error) {
	tmp, err := t.formattedDate()
	result := tmp.Format(time.RFC850)
	return result, err
}

func (t *Times) TimePrevious(day time.Weekday) (time.Time, error) {
	tmp, _ := t.formattedDate()
	currentDay := t.getWeekDayNumber(tmp.Weekday())
	pastDay := t.getWeekDayNumber(day)
	days := 0

	if currentDay > pastDay {
		days = currentDay - pastDay
	} else if currentDay == pastDay {
		days = 7
	} else {
		days = 7 + (currentDay - pastDay)
	}

	return tmp.AddDate(0, 0, -days), nil
}

func (t *Times) TimeNext(day time.Weekday) (time.Time, error) {
	tmp, _ := t.formattedDate()
	currentDay := t.getWeekDayNumber(tmp.Weekday())
	futureDay := t.getWeekDayNumber(day)
	days := 0

	if currentDay < futureDay {
		days = futureDay - currentDay
	} else if currentDay == futureDay {
		days = 7
	} else {
		days = 7 + (futureDay - currentDay)
	}

	return tmp.AddDate(0, 0, days), nil
}

func (t *Times) EndOfHour() (time.Time, error) {
	t2 := *t
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()

	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfMinute() (time.Time, error) {
	t2 := *t
	t.second = 59
	tmp, err := t.formattedDate()

	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfDay() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfWorkWeek() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()

	diff := 5 - t.getWeekDayNumber(tmp.Weekday())

	tmp = tmp.AddDate(0, 0, diff)
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfWeek() (time.Time, error) {
	t2 := *t
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()

	diff := 7 - t.getWeekDayNumber(tmp.Weekday())

	tmp = tmp.AddDate(0, 0, diff)
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfMonth() (time.Time, error) {
	t2 := *t
	t.day = t.maxDaysInMonth()
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfYear() (time.Time, error) {
	t2 := *t
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfDecade() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:3] + "9"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) EndOfCentury() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:2] + "99"
	t.month = 12
	t.day = 31
	t.hour = 23
	t.minute = 59
	t.second = 59
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfCentury() (time.Time, error) {
	t2 := *t
	tmp2 := as.ToString(t.year)
	tmp2 = tmp2[0:2] + "00"
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = as.ToInt(tmp2)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfDecade() (time.Time, error) {
	t2 := *t
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	t.year = t.year - (t.year % 10)
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfYear() (time.Time, error) {
	t2 := *t
	t.month = 01
	t.day = 01
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfDay() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, err := t.formattedDate()
	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfWorkWeek() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()

	diff := t.getWeekDayNumber(tmp.Weekday()) - 1

	tmp = tmp.AddDate(0, 0, -diff)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) StartOfWeek() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()

	diff := t.getWeekDayNumber(tmp.Weekday()) - 1

	tmp = tmp.AddDate(0, 0, -diff)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) StartOfMonth() (time.Time, error) {
	t2 := *t
	t.hour = 00
	t.minute = 00
	t.second = 00
	tmp, _ := t.formattedDate()
	tmp = tmp.AddDate(0, 0, -(t.day)+1)
	if t.reset {
		*t = t2
	}
	return tmp, nil
}

func (t *Times) StartOfHour() (time.Time, error) {
	t2 := *t
	t.minute = 00
	t.second = 00
	tmp, err := t.formattedDate()

	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) StartOfMinute() (time.Time, error) {
	t2 := *t
	t.second = 00
	tmp, err := t.formattedDate()

	if t.reset {
		*t = t2
	}
	return tmp, err
}

func (t *Times) IsWeekend() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Saturday || ti.Weekday() == time.Sunday {
		return true
	}
	return false
}

func (t *Times) IsWorkday() bool {
	if t.IsWeekend() {
		return false
	}
	return true
}

func (t *Times) IsMonday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Monday {
		return true
	}
	return false
}

func (t *Times) IsTuesday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Tuesday {
		return true
	}
	return false
}

func (t *Times) IsWednesday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Wednesday {
		return true
	}
	return false
}

func (t *Times) IsThursday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Thursday {
		return true
	}
	return false
}

func (t *Times) IsFriday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Friday {
		return true
	}
	return false
}

func (t *Times) IsSaturday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Saturday {
		return true
	}
	return false
}

func (t *Times) IsSunday() bool {
	ti, _ := t.formattedDate()
	if ti.Weekday() == time.Sunday {
		return true
	}
	return false
}

func (t *Times) IsLeapYear() bool {
	return t.year%4 == 0 && (t.year%100 != 0 || t.year%400 == 0)
}

func internalIsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (t *Times) AddYear() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(1, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(year, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 1, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, month, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, 7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, week*7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, 1)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, day)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(hour) * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(minute) * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(1 * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) AddSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(second) * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubYear() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-1, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubYears(year int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(-year, 0, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMonth() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, -1, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMonths(month int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, -month, 0)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubWeek() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -7)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubWeeks(week int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -(week * 7))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDay() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -1)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubDays(day int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.AddDate(0, 0, -day)
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubHour() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubHours(hour int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-hour) * time.Hour))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMinute() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubMinutes(minute int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-minute) * time.Minute))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubSecond() (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(-1 * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}

func (t *Times) SubSeconds(second int) (time.Time, error) {
	tmp, err := t.formattedDate()
	if err == nil {
		tmp = tmp.Add(time.Duration(time.Duration(-second) * time.Second))
	}
	if !t.reset {
		t.updateValues(tmp)
	}
	return tmp, err
}
