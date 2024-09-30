package crud

// User represents a user in the database

type User struct {
	ID    uint   gorm:"primaryKey"
	Name  string gorm:"size:100;not null"
	Email string gorm:"size:100;unique;not null"
}