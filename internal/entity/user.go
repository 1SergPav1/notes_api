package entity

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unigue; not null"`
	Password string `grom:"not null"`
}
