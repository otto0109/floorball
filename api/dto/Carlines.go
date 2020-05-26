package dto

type VicciCarlineResult struct {
	Carlines []Carline
}

type Carline struct {
	Name string
	Code string
}
