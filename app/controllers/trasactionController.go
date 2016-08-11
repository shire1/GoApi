package controllers
import (
	"github.com/revel/revel"
	"github.com/shireAli/GoApi/app/models"
	"github.com/shireAli/GoApi/app"
	"log"
	"github.com/shireAli/GoApi/app/util"
	"github.com/shireAli/GoApi/app/encoders"

)
type TransactionController struct  {

			*revel.Controller
}

func (c TransactionController) Get() revel.Result  {
	var transactions []models.Transaction
	if err := app.Db.Where(&transactions).Find(&transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("can't find a Money Info "))
	}
	return c.RenderJson(util.ResponseSuccess(transactions))
}
func (c TransactionController) Find() revel.Result  {
	var transactions models.Transaction
	var id string
	c.Params.Bind(&id , "id")
	transactions.AccNo= id
	if err := app.Db.Where(transactions).Find(&transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money can't find"))
	}
	return c.RenderJson(util.ResponseSuccess(transactions))
}

func (c TransactionController) FindUser() revel.Result  {
	var transactions[] models.Transaction
	var id string
	c.Params.Bind(&id , "id")
	if err := app.Db.Where(models.Transaction{AccNo:id}).Find(&transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money can't find"))
	}
	return c.RenderJson(util.ResponseSuccess(transactions))
}


func (c TransactionController) Create() revel.Result {
var transactions = encoders.EncodeTrasaction(c.Request.Body);

//	if transactions.AccNo=="" {
//		return c.RenderJson(util.ResponseError("transaction failed  customer not registered "))
//	}
	var money models.Money

	money.AccNo = transactions.AccNo

if (transactions.Credit == 0){
	if err := app.Db.Where(&money).First(&money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money finding failed "))
	}
	if (transactions.Debit>money.Balance){
		return c.RenderJson(util.ResponseError("you don't enough balance to debit "))
	} else {

		money.Balance =  money.Balance -transactions.Debit
		money.AccNo = transactions.AccNo
		if err := app.Db.Create(&transactions).Error; err != nil {
			log.Println(err)
			return c.RenderJson(util.ResponseError("Transaction Creation Failed "))
		}
		if err := app.Db.Model(&money).Updates(&money).Error; err != nil {
			log.Println(err)
			return c.RenderJson(util.ResponseError("money Creation Failed pls check "))
		}
	}

}else {

	if err := app.Db.Create(&transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Transaction Credid Creation Failed "))
	}

	money.AccNo = transactions.AccNo
	if founded := app.Db.Where(&money).First(&money).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Upadtes finding Failed"))

	}
	money.Balance=transactions.Credit + money.Balance
	if err := app.Db.Model(&money).Updates(&money).Error; err != nil {
		return c.RenderJson(util.ResponseError("money balance update Failed"))
	}
}

	return c.RenderJson(util.ResponseSuccess(transactions))
}

func (c TransactionController) Delete() revel.Result  {

	var (
		id int
		transaction models.Transaction
	)
	c.Params.Bind(&id , "id")
	if founded := app.Db.First(&transaction , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Delete finding Failed"))

	}
	if err:= app.Db.Delete(&transaction).Error; err != nil {
		return c.RenderJson(util.ResponseError("Delete Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(transaction))
}

