package dto

type Training struct {
	ID      int64  `json:"id"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Gym     string `json:"gym"`
	GymMaps string `json:"gymMaps"`
	Day     string `json:"day"`
}
