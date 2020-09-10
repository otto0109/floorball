package dto

type Player struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Position   string `json:"position"`
	Qoute      string `json:"qoute"`
	Instagram  string `json:"instagram"`
	PictureUrl string `json:"pictureUrl"`
}
