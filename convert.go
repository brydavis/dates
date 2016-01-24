package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t1 := "11/20/1982 20:07:34.830"
	t2 := Convert(t1, "01/02/2006 15:04:05.000", 110)
	fmt.Println(t2)
}

func Convert(date, format string, code int) string {
	t, _ := time.Parse(format, date)
	switch code {
	case 0, 100:
		return t.String()
	case 1:
		return t.Format("01/02/06")
	case 101:
		return t.Format("01/02/2006")
	case 2:
		return t.Format("06.01.02")
	case 102:
		return t.Format("2006.01.02")
	case 3:
		return t.Format("02/01/06")
	case 103:
		return t.Format("02/01/2006")
	case 4:
		return t.Format("02.01.06")
	case 104:
		return t.Format("02.01.2006")
	case 5:
		return t.Format("02-01-06")
	case 105:
		return t.Format("02-01-2006")
	case 6:
		s := t.Format("02 Jan 06") // lowercase needed for month
		return strings.ToLower(s)
	case 106:
		s := t.Format("02 Jan 2006") // lowercase needed for month
		return strings.ToLower(s)
	case 7:
		return t.Format("Jan 02, 06")
	case 107:
		return t.Format("Jan 02, 2006")
	case 108:
		return t.Format("15:04:05")
	case 9, 109:
		s := t.Format("Jan 02 2006 03:04:05.000PM") // lowercase needed for month
		return fmt.Sprintf("%s%s", strings.ToLower(string(s[0])), string(s[1:]))
	case 10:
		return t.Format("06-01-02")
	case 110:
		return t.Format("2006-01-02")
	case 11:
		return t.Format("06/01/02")
	case 111:
		return t.Format("2006/01/02")
	case 12:
		return t.Format("060102")
	case 112:
		return t.Format("20060102")
	case 13, 113:
		s := t.Format("02 Jan 2006 15:04:05.000") // lowercase needed for month
		return fmt.Sprintf("%s%s%s", string(s[:3]), strings.ToLower(string(s[3])), string(s[4:]))
	case 14, 114:
		return t.Format("15:04:05.000")
	case 20, 120:
		return t.Format("2006-01-02 15:04:05")
	case 21, 121:
		return t.Format("2006-01-02 15:04:05.000")
	case 126:
		return t.Format("2006-01-02T15:04:05.000")
	case 127:
		return t.Format("2006-01-02T15:04:05.000Z")
	case 130:
		s := t.Format("02 Jan 2006 03:04:05.000PM") // lowercase needed for month
		return fmt.Sprintf("%s%s%s", string(s[:3]), strings.ToLower(string(s[3])), string(s[4:]))
	case 131:
		s := t.Format("02/01/06 03:04:05.000PM") // lowercase needed for month
		return fmt.Sprintf("%s%s%s", string(s[:3]), strings.ToLower(string(s[3])), string(s[4:]))
	}
	return t.String()
}
