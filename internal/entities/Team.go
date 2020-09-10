package entities

type Team struct {
	ID           int64  `gorm:"primary_key;auto_increment:true"`
	Name         string `gorm:"unique"`
	TeamFoto     string
	PlayerTeam   []PlayerTeam   `gorm:"ForeignKey:team_id;AssociationForeignKey:id"`
	Player       []Player       `gorm:"-"`
	TeamTraining []TeamTraining `gorm:"ForeignKey:team_id;AssociationForeignKey:id"`
	Training     []Training     `gorm:"-"`
}
