package controllers
import (
	"github.com/revel/revel"
	"github.com/shireAli/GoApi/app/models"
	"github.com/shireAli/GoApi/app"
	"log"
	"github.com/shireAli/GoApi/app/util"
	"github.com/shireAli/GoApi/app/encoders"
)
type CustomersController struct  {

			*revel.Controller
}

func (c CustomersController) Get() revel.Result  {
	var customer[] models.Customer

	if err := app.Db.Where(&customer).Find(&customer).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("can't find a Customers "))
	}

	for i ,customers := range customer {
		app.Db.First(&customer[i].Balance , customers.ID)
	}
	return c.RenderJson( util.ResponseSuccess(customer))
}
func (c CustomersController) Find() revel.Result  {
	var customer models.Customer
	var id int
	c.Params.Bind(&id , "id")
	if err := app.Db.First(&customer,id).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Customer can't find"))
	}

		app.Db.First(&customer.Balance , customer.ID)

	return c.RenderJson(util.ResponseSuccess(customer))
}


func (c CustomersController) Create() revel.Result {
var customer = encoders.EncodeCustomer(c.Request.Body);
		if customer.Name==""|| customer.AccNo == "" {
			return c.RenderJson(util.ResponseError("Some required info is missing"));
		}
	if err := app.Db.Create(&customer).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Customer Creation Failed pls check "))
	}
	var Account models.Money
		Account.AccNo = customer.AccNo
	if err := app.Db.Create(&Account).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Customer Account Creation Failed pls check "))
	}

	return c.RenderJson(util.ResponseSuccess(customer))
}

func (c CustomersController) Update() revel.Result  {
	var update = encoders.EncodeCustomer(c.Request.Body);
	var id int
	c.Params.Bind(&id , "id")
	var customer models.Customer
	if founded := app.Db.First(&customer , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Upadtes finding Failed"))

	}
	if err := app.Db.Model(&customer).Updates(&update).Error; err != nil {
		return c.RenderJson(util.ResponseError("Upadtes Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(customer))
}


func (c CustomersController) Delete() revel.Result {

	var (
		id string
		customer models.Customer
	)
	c.Params.Bind(&id, "id")
	customer.AccNo = id
	if err := app.Db.Where(&customer).First(&customer).Delete(customer).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("cusotmer finding failed "))
	}
	var money models.Money
	money.AccNo = customer.AccNo
	if err := app.Db.Where(&money).First(&money).Delete(money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError(" there is an error in many deleting "))
	}
	var transactions models.Transaction
	transactions.AccNo = customer.AccNo
	if err := app.Db.Where(&transactions).Find(&transactions).Delete(transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError(" there is an error in many deleting "))
	}
	return c.RenderJson(util.ResponseSuccess(customer))

}


func (c CustomersController) AndroidDelete() revel.Result  {

	var (
		id string
		customer models.Customer
	)
	c.Params.Bind(&id , "id")
	if err := app.Db.Where(&customer ,id).First(&customer).Delete(customer).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("cusotmer finding failed "))
	}
	var money models.Money
	money.AccNo = customer.AccNo
	if err := app.Db.Where(&money).First(&money).Delete(money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError(" there is an error in many deleting "))
	}
	var transactions models.Transaction
	transactions.AccNo = customer.AccNo
	if err := app.Db.Where(&transactions).Find(&transactions).Delete(transactions).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError(" there is an error in Trasaction deleting "))
	}
	return c.RenderJson(util.ResponseSuccess(customer))
}

