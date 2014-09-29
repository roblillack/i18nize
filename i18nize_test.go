package i18nize

import (
	"testing"
	"time"
)

var liasBirthday time.Time

func init() {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic("Unable to find Berlin!")
	}
	liasBirthday = time.Date(2008, 03, 16, 11, 27, 32, 516823, berlin)
}

func Test_LayoutTime(t *testing.T) {
	if l := LayoutTime("2 Jan", "de"); l != "2. Jan" {
		t.Errorf("Layout ‘%s’ not expected", l)
	}
}

func Test_FormatTime(t *testing.T) {	
	if l := FormatTime(liasBirthday, "de", LayoutTime("2 Jan", "de")); l != "16. Mär" {
		t.Errorf("Formatted time ‘%s’ not expected", l)
	}
	if l := FormatTime(liasBirthday, "de", "am 2. January 2006 um 15:04 Uhr"); l != "am 16. März 2008 um 11:27 Uhr" {
		t.Errorf("Formatted time ‘%s’ not expected", l)
	}
}