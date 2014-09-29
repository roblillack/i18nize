package i18nize

import (
	"time"
	"strings"
)

var shortMonths = map[string]map[time.Month]string {
	"de": {
		time.January: "Jan",
		time.February: "Feb",
		time.March: "Mär",
		time.April: "Apr",
		time.May: "Mai",
		time.June: "Jun",
		time.July: "Jul",
		time.August: "Aug",
		time.September: "Sep",
		time.October: "Okt",
		time.November: "Nov",
		time.December: "Dez",
	},
}

var longMonths = map[string]map[time.Month]string {
	"de": {
		time.January: "Januar",
		time.February: "Februar",
		time.March: "März",
		time.April: "April",
		time.May: "Mai",
		time.June: "Juni",
		time.July: "Juli",
		time.August: "August",
		time.September: "September",
		time.October: "Oktober",
		time.November: "November",
		time.December: "Dezember",
	},
}

var layouts = map[string]map[string]string {
	"de": {
		"2 Jan": "2. Jan",
		"January 2": "2. Januar",
	},
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

func LayoutTime(layout, lang string) string {
	if l := layouts[lang][layout]; l != "" {
		return l
	}
	
	return layout
}

func FormatTime(t time.Time, lang string, layout string) string {
	var suffix = layout
	var buf = make([]string, 0)

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
					if !startsWithLowerCase(suffix[i + 3:]) {
						buf = append(buf, t.Format(suffix[0:i]), formatShortMonth(t, lang))
						suffix = suffix[i+3:]
						break
					}
				}
			// TODO: Parse weekday
			}
		}
		
		buf = append(buf, t.Format(suffix))
		break
	}

	return strings.Join(buf, "")
}
