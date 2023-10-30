package timeutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/BESTSELLER/nightscaler/logger"
)

type Day struct {
	Day   string
	Index int
}

var WeekDay = map[string]Day{
	"Mon": {"Monday", 0},
	"Tue": {"Tuesday", 1},
	"Wed": {"Wednesday", 2},
	"Thu": {"Thursday", 3},
	"Fri": {"Friday", 4},
	"Sat": {"Saturday", 5},
	"Sun": {"Sunday", 6},
}

func GetUptimes(s string) (time.Time, time.Time, error) {
	_, startTime, endTime, err := isUptimeNow(s)

	return startTime, endTime, err
}

func isUptimeNow(s string) (bool, time.Time, time.Time, error) {
	now := time.Now()

	uptimeSplit := strings.Split(s, " ")

	if len(uptimeSplit) != 3 {
		return false, time.Time{}, time.Time{}, fmt.Errorf("'%s' is an invalid uptime format", s)
	}

	weekDays := uptimeSplit[0]
	timeRange := uptimeSplit[1]
	timeZone := uptimeSplit[2]

	days, err := getWeekDays(weekDays)
	if err != nil {
		return false, time.Time{}, time.Time{}, err
	}

	isIt := isCurrentDay(days, now)
	if !isIt {
		return false, time.Time{}, time.Time{}, nil
	}

	startTime, endTime, err := getTimeBoundary(timeRange, timeZone, now)
	if err != nil {
		return false, time.Time{}, time.Time{}, err
	}

	return now.After(startTime) && now.Before(endTime), startTime, endTime, nil
}

func getWeekDays(weekDays string) ([]Day, error) {
	daysFound := []Day{}

	splitDays := strings.Split(weekDays, "-")

	if len(splitDays) == 1 {
		dd, err := getWeekDayIndex(splitDays[0])
		if err != nil {
			return []Day{}, err
		}
		daysFound = append(daysFound, dd)
	} else {
		tmpDay := []Day{}
		for _, b := range splitDays {
			d, err := getWeekDayIndex(b)
			if err != nil {
				return []Day{}, err
			}
			tmpDay = append(tmpDay, d)
		}

		for _, c := range WeekDay {
			if c.Index >= tmpDay[0].Index && c.Index <= tmpDay[1].Index {
				daysFound = append(daysFound, c)
			}
		}

	}

	logger.Log.Debug().Msgf("Days found: %v", daysFound)

	return daysFound, nil
}

func isCurrentDay(days []Day, now time.Time) bool {
	nowAsWeekDay := WeekDay[now.Weekday().String()[0:3]]

	for _, d := range days {
		if d.Index == nowAsWeekDay.Index {
			return true
		}
	}

	return false
}

func getTimeBoundary(timeBoundary string, timeZone string, now time.Time) (time.Time, time.Time, error) {
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	timeBoundarySplit := strings.Split(timeBoundary, "-")

	startTime := getTime(timeBoundarySplit[0], location, now)
	endTime := getTime(timeBoundarySplit[1], location, now)

	return startTime, endTime, err
}

func getTime(s string, location *time.Location, now time.Time) time.Time {
	ss := strings.Split(s, ":")
	hh, _ := strconv.Atoi(ss[0])
	mm, _ := strconv.Atoi(ss[1])

	t := time.Date(now.Year(), now.Month(), now.Day(), hh, mm, 0, 0, location)

	return t
}

func getWeekDayIndex(s string) (Day, error) {
	day, ok := WeekDay[s]

	if !ok {
		return day, fmt.Errorf("'%s' is an invalid day", s)
	}

	return day, nil
}
