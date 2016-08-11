package models

type User struct {
	ID 				int64  		`gorm:"AUTO_INCREMENT, primary_key" json:"id"`
	Name 			string  	`json:"name"`
	Email 			string    	`json:"email"`
	Password		string    	`json:"password"`

}
