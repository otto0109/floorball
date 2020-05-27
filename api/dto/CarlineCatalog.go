package dto

type VicciCarlineCatalogResult struct {
	Carline     CarlineCatalog
	Salesgroups []Salesgroup
}

type CarlineCatalog struct {
	Name        string
	Code        string
	Salesgroups []Salesgroup
}

type Salesgroup struct {
	Name   string
	Code   string
	Models []Model
}

type ModelResult struct {
	Models []Model
}

type Model struct {
	Name      string
	Code      string
	Version   string
	ModelYear string `json:"year"`
}
