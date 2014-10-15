package i18nize

import (
	"strings"
	"time"
)

var shortDays = map[string]map[time.Weekday]string{
	"de": {
		time.Monday:    "Mo",
		time.Tuesday:   "Di",
		time.Wednesday: "Mi",
		time.Thursday:  "Do",
		time.Friday:    "Fr",
		time.Saturday:  "Sa",
		time.Sunday:    "So",
	},
}

var longDays = map[string]map[time.Weekday]string{
	"de": {
		time.Monday:    "Montag",
		time.Tuesday:   "Dienstag",
		time.Wednesday: "Mittwoch",
		time.Thursday:  "Donnerstag",
		time.Friday:    "Freitag",
		time.Saturday:  "Samstag",
		time.Sunday:    "Sonntag",
	},
}

var shortMonths = map[string]map[time.Month]string{
	"de": {
		time.January:   "Jan",
		time.February:  "Feb",
		time.March:     "Mär",
		time.April:     "Apr",
		time.May:       "Mai",
		time.June:      "Jun",
		time.July:      "Jul",
		time.August:    "Aug",
		time.September: "Sep",
		time.October:   "Okt",
		time.November:  "Nov",
		time.December:  "Dez",
	},
}

var longMonths = map[string]map[time.Month]string{
	"de": {
		time.January:   "Januar",
		time.February:  "Februar",
		time.March:     "März",
		time.April:     "April",
		time.May:       "Mai",
		time.June:      "Juni",
		time.July:      "Juli",
		time.August:    "August",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Dezember",
	},
}

var layouts = map[string]map[string]string{
	"de": {
		"2 Jan":          "2. Jan",
		"2 January":      "2. January",
		"2 Jan 06":       "2. Jan 06",
		"2 Jan 2006":     "2. Jan 2006",
		"2 January 2006": "2. January 2006",
		"Jan 2006":       "Jan 2006",
		"January 2006":   "January 2006",
		"Mon 2 Jan":      "Mon, 2. Jan",
	},
}

func formatShortDay(t time.Time, lang string) string {
	if m := shortDays[lang][t.Weekday()]; m != "" {
		return m
	}

	return t.Format("Mon")
}

func formatLongDay(t time.Time, lang string) string {
	if m := longDays[lang][t.Weekday()]; m != "" {
		return m
	}

	return t.Format("Monday")
}

func formatShortMonth(t time.Time, lang string) string {
	if m := shortMonths[lang][t.Month()]; m != "" {
		return m
	}

	return t.Format("Jan")
}

func formatLongMonth(t time.Time, lang string) string {
	if m := longMonths[lang][t.Month()]; m != "" {
		return m
	}

	return t.Format("January")
}

func startsWithLowerCase(str string) bool {
	if len(str) == 0 {
		return false
	}

	c := str[0]
	return 'a' <= c && c <= 'z'
}

func getLayout(layout, lang string) string {
	if len(lang) < 2 {
		return layout
	}

	if l := layouts[lang[:2]][layout]; l != "" {
		return l
	}

	return layout
}

// Formats the time instant t for language lang using layout,
// while changing language dependent characteristics of the layout
// where applicable.
func LayoutTime(t time.Time, lang string, layout string) string {
	return FormatTime(t, lang, getLayout(layout, lang))
}

// Formats the time instant t for language lang using layout.
// This function is the i18nized equivalent to time's time.Format.
func FormatTime(t time.Time, lang string, layout string) string {
	var suffix = layout
	var buf = make([]string, 0)

	if len(lang) < 2 {
		return t.Format(layout)
	}
	lang = lang[:2]

	for suffix != "" {
		for i := 0; i < len(suffix); i++ {
			switch c := int(suffix[i]); c {
			case 'J': // January, Jan
				if len(suffix) >= i+3 && suffix[i:i+3] == "Jan" {
					if len(suffix) >= i+7 && suffix[i:i+7] == "January" {
						buf = append(buf, t.Format(suffix[0:i]), formatLongMonth(t, lang))
						suffix = suffix[i+7:]
						break
					}
					if !startsWithLowerCase(suffix[i+3:]) {
						buf = append(buf, t.Format(suffix[0:i]), formatShortMonth(t, lang))
						suffix = suffix[i+3:]
						break
					}
				}
			case 'M': // Monday, Mon
				if len(suffix) >= i+3 {
					if suffix[i:i+3] == "Mon" {
						if len(suffix) >= i+6 && suffix[i:i+6] == "Monday" {
							buf = append(buf, t.Format(suffix[0:i]), formatLongDay(t, lang))
							suffix = suffix[i+6:]
							break
						}
						if !startsWithLowerCase(suffix[i+3:]) {
							buf = append(buf, t.Format(suffix[0:i]), formatShortDay(t, lang))
							suffix = suffix[i+3:]
							break
						}
					}
				}
			}
		}

		buf = append(buf, t.Format(suffix))
		break
	}

	return strings.Join(buf, "")
}
