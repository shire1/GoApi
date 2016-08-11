package controllers
import (
	"github.com/revel/revel"
	"github.com/shireAli/GoApi/app/models"
"github.com/shireAli/GoApi/app"
	"log"
	"github.com/shireAli/GoApi/app/util"
	"github.com/shireAli/GoApi/app/encoders"
)
type MoneyController struct  {

			*revel.Controller
}

func (c MoneyController) Get() revel.Result  {
	var money []models.Money
	if err := app.Db.Where(&money).Find(&money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("can't find a Money Info "))
	}
	return c.RenderJson(util.ResponseSuccess(money))
}
func (c MoneyController) Find() revel.Result  {
	var money models.Money
	var id string
	c.Params.Bind(&id , "id")
	money.AccNo = id
	if err := app.Db.First(&money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money can't find"))
	}
	return c.RenderJson(util.ResponseSuccess(money))
}


func (c MoneyController) Create() revel.Result {
var money = encoders.EncodeMoney(c.Request.Body);
		if money.AccNo== ""|| money.Balance == 0{
			return c.RenderJson(util.ResponseError("Some required info is missing"));
		}


	if err := app.Db.Create(&money).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money Creation Failed pls check "))
	}

	return c.RenderJson(util.ResponseSuccess(money))
}

func (c MoneyController) Update() revel.Result  {
	var update = encoders.EncodeMoney(c.Request.Body);
	var id int
	c.Params.Bind(&id , "id")
	var money models.Money
	if founded := app.Db.First(&money, id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Upadtes finding Failed"))

	}
	if err := app.Db.Model(&money).Updates(&update).Error; err != nil {
		return c.RenderJson(util.ResponseError("Upadtes Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(money))
}


func (c MoneyController) Delete() revel.Result  {

	var (
		id int
		money models.Money
	)
	c.Params.Bind(&id , "id")
	if founded := app.Db.First(&money , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Delete finding Failed"))

	}
	if err:= app.Db.Delete(&money).Error; err != nil {
		return c.RenderJson(util.ResponseError("Delete Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(money))
}

