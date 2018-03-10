package app_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/swallowstalker/cutibot-importer/app"
)

func TestParsingValidInputJSON(t *testing.T) {

	fp, err := os.Open("../tests/valid_input.json")
	if err != nil {
		t.Errorf("Error while opening file pointer: " + err.Error())
	}

	var holidayList []app.Holiday
	decoder := json.NewDecoder(fp)
	err = decoder.Decode(&holidayList)
	if err != nil {
		t.Errorf("Error while decoding data from file: " + err.Error())
	}

	parsedTime := holidayList[0].DateList[0].Time

	if len(holidayList[0].DateList) > 1 {
		t.Errorf("Data is incorrect, date list is more than 1")
	} else if parsedTime.Format(app.DATE_STRING_LAYOUT) != "2018-01-01" {
		t.Errorf("Data is incorrect, " + parsedTime.String())
	}

}

func TestParsingInvalidInputJSON(t *testing.T) {

	fp, err := os.Open("../tests/invalid_input.json")
	if err != nil {
		t.Errorf("Error while opening file pointer: " + err.Error())
	}

	var holidayList []app.Holiday
	decoder := json.NewDecoder(fp)
	err = decoder.Decode(&holidayList)
	if err != nil {
		t.Errorf("Error while decoding data from file: " + err.Error())
	}

	parsedTime := holidayList[0].DateList[0].Time

	if len(holidayList[0].DateList) > 1 {
		t.Errorf("Data is incorrect, date list is more than 1")
	} else if parsedTime.Format(app.DATE_STRING_LAYOUT) != "2018-01-03" {
		t.Errorf("Data is incorrect, " + parsedTime.String())
	}

}
