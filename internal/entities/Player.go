package entities

type Player struct {
	ID         int64 `gorm:"primary_key;"`
	Name       string
	Surname    string
	Position   string
	Qoute      string
	Instagram  string
	PictureUrl string
	PlayerTeam []PlayerTeam `gorm:"ForeignKey:player_id;AssociationForeignKey:id"`
}
