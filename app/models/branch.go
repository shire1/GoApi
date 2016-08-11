package models
import "time"

type Branch struct {
	ID 					int64  		`gorm:"AUTO_INCREMENT, primary_key" json:"id"`
	BranchName 			string  	`json:"branchName"`
	AdminName 			string  	`json:"adminName"`
	CreatedAt      	 	time.Time   `json:"createdAt"`
	UpdatedAt       	time.Time   `json:"updatedAt"`
	DeletedAt      		*time.Time 	`sql:"index",json:"deletedAt"`

}
