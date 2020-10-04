package dao

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CurrentDate struct {
	Month         string
	Day           int
	Year          int
	FormattedDate string
}

// Get the current date ...
func (cd *CurrentDate) GetCurrentDate() string {
	d := time.Now()
	year, month, day := d.Date()
	return fmt.Sprintf("%d-%s-%d", day, month, year)
}

// Set the current date ...
func (cd *CurrentDate) SetDate() {
	d := time.Now()
	year, month, day := d.Date()
	cd.Day = day
	cd.Month = month.String()
	cd.Year = year
	cd.FormattedDate = fmt.Sprintf("%d-%s-%d", day, month, year)
}

// Set the given date as current date ...
func (cd *CurrentDate) SetGivenDate(fmtDt string) {
	cd.FormattedDate = fmtDt
	params := strings.Split(fmtDt, "-")
	cd.Day, _ = strconv.Atoi(params[0])
	cd.Month = params[1]
	cd.Year, _ = strconv.Atoi(params[2])
}
