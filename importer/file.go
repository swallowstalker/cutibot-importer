package importer

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/swallowstalker/cutibot-importer/app"
)

type FileImporter struct {
}

func (fi FileImporter) Import(filename string) ([]app.Holiday, error) {

	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	jsonData, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	var holiday []app.Holiday
	err = json.Unmarshal(jsonData, &holiday)
	if err != nil {
		return nil, err
	}

	return holiday, nil
}
