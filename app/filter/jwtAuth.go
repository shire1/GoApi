package filter
import (
	"github.com/revel/revel"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"strconv"
	"github.com/shireAli/GoApi/app/util"
)


type JwtAuth struct  {
	*revel.Controller

}
func (c JwtAuth) CheckToken() revel.Result  {
	var userWithToken = c.Request.Header.Get("token")

	token, err := jwt.Parse(userWithToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		appSecret, _:= revel.Config.String("app.secret")
		return []byte(appSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			c.Session["email"] = claims["email"].(string);
			c.Session["id"] = strconv.Itoa(int(claims["id"].(float64)));
			return nil
		}
	} else {
		return c.RenderJson(util.ResponseError("Sory invalied key"))
	}
	return c.RenderJson(util.ResponseSuccess(token))
}
func init() {
	revel.InterceptMethod(JwtAuth.CheckToken , revel.BEFORE)
}