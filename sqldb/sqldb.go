package sqldb

import (
	// "time"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

// type Reservation struct {
//     gorm.Model
//     StartTime time.Time
//     UserSlackID string
//     RoomID string
// }
//
// type User struct {
//     gorm.Model
//     SlackID string
//     Reservations []Reservation `gorm:"foreignkey:UserSlackID;association_foreignkey:SlackID"`
// }

// SetupDB is a function that connects to a gorm db
func SetupDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	// defer db.Close()

	// db.AutoMigrate(&Reservation{}, &User{})

	return db
}
