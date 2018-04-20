package ctime

/*
#include <time.h>
*/
import "C"
import (
	"time"
)

// Strftime : 'C' style time formatter
// exzmple usage: fmt.Println(strftime(time.Now(), "%Y-%m-%d %H:%M:%S"))
func Strftime(time time.Time, format string) string {

	const buflen = 64
	var tm C.struct_tm
	buffer := make([]C.char, buflen)
	tm.tm_sec = C.int(time.Second())
	tm.tm_min = C.int(time.Minute())
	tm.tm_hour = C.int(time.Hour())
	tm.tm_mday = C.int(time.Day())
	tm.tm_mon = C.int(time.Month())
	tm.tm_wday = C.int(time.Weekday())
	tm.tm_yday = C.int(time.YearDay())
	tm.tm_year = C.int(time.Year() - 1900)
	zone, _ := time.Zone()
	tm.tm_zone = C.CString(zone)

	C.strftime(&buffer[0], buflen, C.CString(format), &tm)

	return C.GoString(&buffer[0])
}
