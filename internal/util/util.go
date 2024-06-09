package util

import (
	"strings"
	"time"
)

func TimeToFractional(t time.Time) float64 {
	f := float64(t.UnixNano()) / 1e9
	return f
}

func MatchWithUpper(value, search string) bool {
	if value == search {
		return true
	}
	parts := strings.Split(search, " ")
	upper := strings.ToUpper(strings.Join(parts, "_"))
	return value == upper
}
