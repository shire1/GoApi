package encoders
import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/shireAli/GoApi/app/models"
)
func EncodeSingleUser(body io.ReadCloser) (user *models.User) {
	data , _:= ioutil.ReadAll(body);
	if err :=  json.Unmarshal(data , &user); err != nil {
		log.Println("there is a problem on data binding")
	}
	 return user
}
