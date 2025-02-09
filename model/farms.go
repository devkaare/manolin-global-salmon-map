package model

type Farms struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	// LocalityNumber string  `json:"localityNumber"`
	// Municipality   string  `json:"municipality"`
	Name          string `json:"name"`
	Organizations string `json:"organizations"`
	// ProductionArea string  `json:"productionArea"`
	Regions string `json:"regions"`
	Species string `json:"species"`
}
