package model

type Farms struct {
	Lat            float32 `json:"lat,omitempty"`
	Lon            float32 `json:"lon,omitempty"`
	LocalityNumber string  `json:"localityNumber,omitempty"`
	Municipality   string  `json:"municipality,omitempty"`
	Name           string  `json:"name,omitempty"`
	Organizations  string  `json:"organizations,omitempty"`
	ProductionArea string  `json:"productionArea,omitempty"`
	Regions        string  `json:"regions,omitempty"`
	Species        string  `json:"species,omitempty"`
}
