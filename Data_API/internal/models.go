package internal

type PropertiesType struct {
	Type string `json:"type"`
}

type Properties struct {
	Address  PropertiesType `json:"address"`
	Id       PropertiesType `json:"id"`
	Location PropertiesType `json:"location"`
	Name     PropertiesType `json:"name"`
	Phone    PropertiesType `json:"phone"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

type Location struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Place struct {
	Address  string   `json:"address"`
	Id       int      `json:"id"`
	Location Location `json:"location"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}
