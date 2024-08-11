package internal

type Shema struct {
	Places Place
}

type Location struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Place struct {
	ID        int      `json:"id"`
	Name      string   `json:"string"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	Locations Location `json:"location"`
}
