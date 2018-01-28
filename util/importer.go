package util

import (
	"encoding/json"
	"time"
)

type Holiday struct {
	Description         string               `json:"description"`
	DateList            []TimeFromDateString `json:"date_list"`
	HolidayStreak       HolidayStreak        `json:"holiday_streak"`
	LeaveRecommendation []TimeFromDateString `json:"leave_recommendation"`
}

type HolidayStreak struct {
	Start TimeFromDateString `json:"start"`
	End   TimeFromDateString `json:"end"`
}

type TimeFromDateString struct {
	Time time.Time
}

const DATE_STRING_LAYOUT = "2006-01-02"

func (d *TimeFromDateString) UnmarshalJSON(data []byte) error {

	var unconvertedTime string

	err := json.Unmarshal(data, &unconvertedTime)
	if err != nil {
		return err
	}

	if unconvertedTime == "" {
		d = nil
		return nil
	}

	parsedTime, err := time.Parse(DATE_STRING_LAYOUT, unconvertedTime)
	if err != nil {
		return err
	}

	d.Time = parsedTime

	return nil
}
