package entities

type Article struct {
	ID          int64 `gorm:"primary_key;"`
	Thumbnail   string
	Title       string
	Article     string
	ArticleTeam []ArticleTeam `gorm:"ForeignKey:team_id;AssociationForeignKey:id"`
	Teams       []Team        `gorm:"-"`
	Category    string
}
