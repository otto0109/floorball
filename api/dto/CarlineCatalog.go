package dto

type CarlineCatalog struct {
	Name        string
	Code        string
	Salesgroups []Salesgroup
}

type Salesgroup struct {
	Name   string
	Code   string
	Models []Model `json:"referenceModels"`
}

type Model struct {
	Name      string
	Code      string
	Version   string
	ModelYear string
}
