package entities

type Team struct {
	ID         int          `gorm:"primary_key;auto_increment:true"`
	Name       string       `gorm:"unique"`
	PlayerTeam []PlayerTeam `gorm:"ForeignKey:team_id;AssociationForeignKey:id"`
	Player     []Player     `gorm:"-"`
}
