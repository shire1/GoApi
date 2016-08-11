package controllers
import (
	"github.com/revel/revel"
	"github.com/shireAli/GoApi/app/models"
"github.com/shireAli/GoApi/app"
	"log"
	"github.com/shireAli/GoApi/app/util"
	"github.com/shireAli/GoApi/app/encoders"
	"github.com/dgrijalva/jwt-go"
)
type UsersController struct  {

			*revel.Controller
}

func (c UsersController) Find() revel.Result  {
	var user models.User
	var id int
	c.Params.Bind(&id , "id")
	if err := app.Db.First(&user,id).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("User not exist"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}
func (c UsersController) FindAll() revel.Result  {
	var user [] models.User

	if founded:= app.Db.Find(&user).RowsAffected; founded<1{

		return c.RenderJson(util.ResponseError("There is no Admin User! Pls Register New"))
	}

	return c.RenderJson(util.ResponseSuccess(user))
}


func (c UsersController) Create() revel.Result {
var user = encoders.EncodeSingleUser(c.Request.Body);
		if user.Email==""|| user.Password == "" {
			return c.RenderJson(util.ResponseError("Some required info is missing"));
		}
	if err := app.Db.Create(&user).Error; err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("User Creation Failed pls check "))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

func (c UsersController) Login() revel.Result  {
	var user = encoders.EncodeSingleUser(c.Request.Body);
	if user.Email=="" || user.Password == ""{
		return c.RenderJson(util.ResponseError("User Login Failed"))
	}
	if userData:= app.Db.Where(&user).First(&user).RowsAffected; userData<1{
		return c.RenderJson(util.ResponseError("User Not founded"))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"id": user.ID,
		"email":user.Email,
		"name":user.ID,
	});
	appSecret, _:= revel.Config.String("app.secret")
	dataToken,err := token.SignedString([]byte(appSecret));
	if err != nil {
		log.Println(err)
	return c.RenderJson(util.ResponseError("Token Genaration Faild"))
	}
	var usertoken models.UserToken
	usertoken.Name = user.Name
	usertoken.Email = user.Email
	usertoken.Token = dataToken
	return c.RenderJson(util.ResponseSuccess(usertoken))
}

func (c UsersController) Update() revel.Result  {
	var update = encoders.EncodeSingleUser(c.Request.Body);
	var id int
	c.Params.Bind(&id , "id")
	var user models.User
	if founded := app.Db.First(&user , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Upadtes finding Failed"))

	}
	if err := app.Db.Model(&user).Updates(&update).Error; err != nil {
		return c.RenderJson(util.ResponseError("Upadtes Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}


func (c UsersController) Delete() revel.Result  {

	var (
		id int
		user models.User
	)
	c.Params.Bind(&id , "id")
	if founded := app.Db.First(&user , id).RowsAffected; founded<1 {
		return c.RenderJson(util.ResponseError("Delete finding Failed"))

	}
	if err:= app.Db.Delete(&user).Error; err != nil {
		return c.RenderJson(util.ResponseError("Delete Proccess Failed"))
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

