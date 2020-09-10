package dto

type Team struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	TeamFoto string     `json:"teamFoto"`
	Player   []Player   `json:"player"`
	Training []Training `json:"training"`
}
