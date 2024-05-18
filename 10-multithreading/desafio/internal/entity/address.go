package entity

type Address struct {
	CEP          string `json:"cep"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Neighborhood string `json:"neighborhood"`
}

func NewAddress(cep, streetName, city, state, neighborhood string) *Address {
	return &Address{
		CEP:          cep,
		StreetName:   streetName,
		City:         city,
		State:        state,
		Neighborhood: neighborhood,
	}
}
