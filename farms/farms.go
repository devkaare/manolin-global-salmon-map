package farms

import (
	"encoding/json"
	"os"

	"github.com/devkaare/manolin-global-salmon-map/model"
)

var filesToRead = []string{"json-files/canada_farms.json", "json-files/chile_farms.json", "json-files/new_brunswick_farms.json", "json-files/new_zealand_farms.json", "json-files/newfoundland_farms.json", "json-files/norway_farms.json", "json-files/scotland_farms.json", "json-files/tasmania_farms.json"}

func Get() ([]model.Farms, error) {
	var farms []model.Farms

	for _, file := range filesToRead {
		rawData, err := os.ReadFile(file)
		if err != nil {
			return farms, err
		}

		data := []model.Farms{}
		if err := json.Unmarshal(rawData, &data); err != nil {
			return farms, err
		}

		farms = append(farms, data...)
	}

	return farms, nil
}
