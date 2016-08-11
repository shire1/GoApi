package models
import "time"

type Transaction struct {
	ID 					int64  		`gorm:"AUTO_INCREMENT, primary_key" json:"id"`
	AccNo 				string 		`json:"accNo"`
	Credit 				int64    	`json:"credit"`
	Debit  				int64    	`json:"debit"`
	BranchName  		string    	`json:"branchName"`
	CreatedAt       	time.Time   `json:"createdAt"`
	UpdatedAt       	time.Time 	`json:"updated"`
	DeletedAt       	*time.Time 	`sql:"index",json:"deletedAt"`

}
