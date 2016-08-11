package models
import "time"

type Money struct {
	ID 					int64  			`gorm:"AUTO_INCREMENT, primary_key" json:"id"`
	AccNo 				string  		`json:"accNo"`
	Balance 			int64    		`json:"balance"`
	Branch 				string    		`json:"branch"`
	CreatedAt       	time.Time  		`json:"createdAt"`
	UpdatedAt       	time.Time 		`json:"updated"`
	DeletedAt       	*time.Time 		`sql:"index",json:"deletedAt"`

}
