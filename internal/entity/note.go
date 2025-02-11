package entity

type Note struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Body   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
