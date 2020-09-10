package entities

type Training struct {
	ID      int64  `gorm:"primary_key;auto_increment:true"`
	Start   string `gorm:"size:5"`
	End     string `gorm:"size:5"`
	Gym     string
	GymMaps string
	Day     string `gorm:"size:2"`
}
