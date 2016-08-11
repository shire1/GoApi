package controllers
import (
	"github.com/revel/revel"
	"github.com/shireAli/GoApi/app/models"
"github.com/shireAli/GoApi/app"
	"log"
	"github.com/shireAli/GoApi/app/util"
	"github.com/shireAli/GoApi/app/encoders"
)
type BranchController struct  {

			*revel.Controller
}

func (c BranchController) Get() revel.Result  {
	var branch []models.Branch
	if err := app.Db.Where(&branch).Find(&branch).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("can't find a Branch Info "))
	}
	return c.RenderJson(util.ResponseSuccess(branch))
}
func (c BranchController) Find() revel.Result  {
	var branch models.Branch
	var id int
	c.Params.Bind(&id , "id")
	if err := app.Db.First(&branch,id).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money can't find"))
	}
	return c.RenderJson(util.ResponseSuccess(branch))
}


func (c BranchController) Create() revel.Result {
var branch = encoders.EncodeBranch(c.Request.Body);

	if err := app.Db.Create(&branch).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("money Creation Failed pls check "))
	}

	return c.RenderJson(util.ResponseSuccess(branch))
}

func (c BranchController) Update() revel.Result  {
	var update = encoders.EncodeBranch(c.Request.Body);
	var id int
	c.Params.Bind(&id , "id")
	var branch models.Branch
	if founded := app.Db.First(&branch, id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Upadtes finding Failed"))

	}
	if err := app.Db.Model(&branch).Updates(&update).Error; err != nil {
		return c.RenderJson(util.ResponseError("Upadtes Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(branch))
}


func (c BranchController) Delete() revel.Result  {

	var (
		id int
		branch models.Branch
	)
	c.Params.Bind(&id , "id")
	if founded := app.Db.First(&branch , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Delete finding Failed"))

	}
	if err:= app.Db.Delete(&branch).Error; err != nil {
		return c.RenderJson(util.ResponseError("Delete Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(branch))
}

