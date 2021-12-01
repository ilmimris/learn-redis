package domain

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
