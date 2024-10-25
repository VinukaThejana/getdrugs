package types

// Medicine is a struct that contains the name and dosage of a medicine
type Medicine struct {
	Name   string  `json:"name"`
	Dosage *string `json:"dosage"`
}

// Prescription is a struct that contains the list of medicines, causes and fatality of a prescription
type Prescription struct {
	Medicine []Medicine `json:"medicine"`
	Causes   []string   `json:"causes"`
	Fatality int        `json:"fatality"`
}
