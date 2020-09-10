package entities

type TeamTraining struct {
	ID         int64 `gorm:"primary_key;auto_increment:true"`
	TeamID     int64
	TrainingID int64
}
