package query

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestGenerateImportScript calls GenerateImportScript
// for a valid return value.
type User struct {
	gorm.Model
	FirstName      string    
	LastName       string    
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func TestGenerateImportScript(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
        t.Fatalf("DB connection error should be nil")
    }
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	newUsers := []*User{}
	newUsers = append(newUsers, 
		&User{
			FirstName: "John",
			LastName: "Doe",},
		&User{
			FirstName: "Jane",
			LastName: "Michael",
		},
	)

    // ... Create a new user record...
    result := db.Create(newUsers)
	if result.Error != nil {
        panic("failed to create user: " + result.Error.Error())
    }
	
	if result.Error != nil {
        panic("failed to create user: " + result.Error.Error())
    }
	
	path, err := GenerateImportScript(db, User{}, "hello", nil)
	fmt.Printf("Pathhhh: %s", path)
	if err == nil {
        t.Fatalf("Error should be nil")
    }
}
