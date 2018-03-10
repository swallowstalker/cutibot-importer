package importer_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/swallowstalker/cutibot-importer/app"
	"github.com/swallowstalker/cutibot-importer/importer"
)

func TestNormalImport(t *testing.T) {
	fi := importer.FileImporter{}
	resultHoliday, err := fi.Import("../tests/valid_input.json")
	if err != nil {
		t.Errorf("valid json input should return no error: %v", err)
	}

	holidayDate, err := time.Parse(app.DATE_STRING_LAYOUT, "2018-01-01")
	if err != nil {
		t.Errorf("error while parsing time: %v", err)
	}

	startDate, err := time.Parse(app.DATE_STRING_LAYOUT, "2017-12-29")
	if err != nil {
		t.Errorf("error while parsing time: %v", err)
	}

	endDate, err := time.Parse(app.DATE_STRING_LAYOUT, "2018-01-02")
	if err != nil {
		t.Errorf("error while parsing time: %v", err)
	}

	expectedHoliday := []app.Holiday{
		app.Holiday{
			Description: "Tahun Baru Masehi",
			DateList: []app.TimeFromDateString{
				app.TimeFromDateString{
					holidayDate,
				},
			},
			HolidayStreak: app.HolidayStreak{
				Start: app.TimeFromDateString{
					Time: startDate,
				},
				End: app.TimeFromDateString{
					Time: endDate,
				},
			},
			LeaveRecommendation: []app.TimeFromDateString{
				app.TimeFromDateString{
					Time: startDate,
				},
				app.TimeFromDateString{
					Time: endDate,
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedHoliday, resultHoliday) {
		t.Errorf("got\n %v\n\n expected\n %v", resultHoliday, expectedHoliday)
	}
}

func TestFileNotFound(t *testing.T) {
	fi := importer.FileImporter{}
	_, err := fi.Import("../tests/valid_input2.json")
	if err == nil {
		t.Errorf("file not found should raise error")
	}
}

func TestInvalidJSON(t *testing.T) {
	fi := importer.FileImporter{}
	_, err := fi.Import("../tests/invalid_input.json")
	if err == nil {
		t.Errorf("invalid json should raise error")
	}
}
