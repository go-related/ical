package ical

import "time"

func newDatePtr(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &t
}

func datePtr(date time.Time) *time.Time {
	return &date
}

func parseLocation(loc string) *time.Location {
	location, _ := time.LoadLocation(loc)
	return location
}
