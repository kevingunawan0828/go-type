package gotype

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"github.com/payfazz/go-errors"
)

// error list
var (
	ErrDateNotValid = errors.New("date is not valid")
)

const (
	// INPUTDATEFORMAT Format of input date
	INPUTDATEFORMAT = "20060102"

	// INPUTTIMEFORMAT Format of input time
	INPUTTIMEFORMAT = "2006-01-02T15:04:05-0700"
)

// IntDate custom type for date in INT form
type IntDate int

// IsValid Check if IntDate is Date
func (i IntDate) IsValid() bool {
	_, err := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	if err != nil {
		return false
	}
	return true
}

//IsLastDayofMonth Return if date is the last day of the month or not.
func (i IntDate) IsLastDayofMonth() bool {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	currentYear, currentMonth, _ := date.Date()
	firstOfMonth, _ := time.Parse(INPUTDATEFORMAT, fmt.Sprintf("%d%02d01", currentYear, currentMonth))

	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	if i.String() == lastOfMonth.Format(INPUTDATEFORMAT) {
		return true
	}
	return false
}

//IsFirstDayofMonth Return if date is the first day of the month or not.
func (i IntDate) IsFirstDayofMonth() bool {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	currentYear, currentMonth, _ := date.Date()
	firstOfMonth, _ := time.Parse(INPUTDATEFORMAT, fmt.Sprintf("%d%d01", currentYear, currentMonth))

	if i.String() == firstOfMonth.Format(INPUTDATEFORMAT) {
		return true
	}
	return false
}

// LastDayofTheMonth return last day of the month.
func (i IntDate) LastDayofTheMonth() IntDate {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	currentYear, currentMonth, _ := date.Date()
	firstDay, _ := time.Parse(INPUTDATEFORMAT, fmt.Sprintf("%d%02d01", currentYear, currentMonth))
	lastDay, _ := strconv.ParseInt(firstDay.AddDate(0, 1, -1).Format(INPUTDATEFORMAT), 10, 64)
	return IntDate(lastDay)
}

// FirstDayofTheMonth return first day of the month.
func (i IntDate) FirstDayofTheMonth() IntDate {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	currentYear, currentMonth, _ := date.Date()
	firstDay, _ := time.Parse(INPUTDATEFORMAT, fmt.Sprintf("%d%02d01", currentYear, currentMonth))
	result, _ := strconv.ParseInt(firstDay.Format(INPUTDATEFORMAT), 10, 64)
	return IntDate(result)
}

// ToTime return Time.
func (i IntDate) ToTime() (*time.Time, error) {
	result, err := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Value Value for IntDate Scan
func (i IntDate) Value() driver.Value {
	return int(i)
}

func (i IntDate) String() string {
	return strconv.Itoa(int(i))
}

//Month Return month of the date.
func (i IntDate) Month() int {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	_, currentMonth, _ := date.Date()
	result, _ := strconv.ParseInt(fmt.Sprintf("%d", currentMonth), 10, 64)
	return int(result)
}

//Year Return year of the date.
func (i IntDate) Year() int {
	date, _ := time.Parse(INPUTDATEFORMAT, strconv.Itoa(int(i)))
	currentYear, _, _ := date.Date()
	return currentYear
}
