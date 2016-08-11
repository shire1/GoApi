package encoders
import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/shireAli/GoApi/app/models"
)
func EncodeTrasaction(body io.ReadCloser) (trasaction *models.Transaction) {
	data , _:= ioutil.ReadAll(body);
	if err :=  json.Unmarshal(data , &trasaction); err != nil {
		log.Println("there is a problem on data binding")
	}
	 return trasaction
}
