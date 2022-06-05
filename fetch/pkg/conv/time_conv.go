package conv

import (
	"strings"
	"time"
)

// ParseDate parse date string to format 2006-01-02
func ParseDate(date string) time.Time {

	// incase date format like this
	// 2022/05/16 19:42:29
	// 2022-05-17 12:12:21
	// 2022-05-21T05:56:40.089Z
	if strings.Contains(date, " ") {
		split := strings.Split(date, " ")
		date = split[0]
	} else if strings.Contains(date, "T") {
		split := strings.Split(date, "T")
		date = split[0]
	}

	var parsedDate time.Time
	// date format still can have different pattern
	// 2022/05/16
	// 2022-05-17
	if strings.Contains(date, "/") {
		parsedDate, _ = time.Parse("2006/01/02", date)
	} else {
		parsedDate, _ = time.Parse("2006-01-02", date)
	}
	return parsedDate
}

// example date
// 2022/05/16 19:42:29
// 2022-05-17 12:12:21
// 2022-05-21T05:56:40.089Z
// 31/5/2022
// Mon May 23 07:03:24 GMT+07:00 2022  // RubyDate
