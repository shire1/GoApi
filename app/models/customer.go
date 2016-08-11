package models
import "time"

type Customer struct {
	ID 				int64  		`gorm:"AUTO_INCREMENT, primary_key" json:"id"`
	Name 			string  	`json:"name"`
	AccNo 			string    	`json:"accNo"`
	Address			string    	`json:"address"`
	Email			string    	`json:"email"`
	Balance 		Money        `json:"money"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	DeletedAt      *time.Time 	`sql:"index",json:"deletedAt"`


}
