package internal

type Schema struct {
	Properties struct {
		Id       Field `json:"id"`
		Address  Field `json:"address"`
		Location Field `json:"location"`
		Name     Field `json:"name"`
		Phone    Field `json:"phone"`
	} `json:"properties"`
}

type Field struct {
	Type string `json:"type"`
}

type Place struct {
	Id       int      `json:"id"`
	Address  string   `json:"address"`
	Location Location `json:"location"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}

type Location struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}
