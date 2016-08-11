package encoders
import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/shireAli/GoApi/app/models"
)
func EncodeMoney(body io.ReadCloser) (money *models.Money) {
	data , _:= ioutil.ReadAll(body);
	if err :=  json.Unmarshal(data , &money); err != nil {
		log.Println("there is a problem on data binding")
	}
	 return money
}
