package encoders
import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/shireAli/GoApi/app/models"
)
func EncodeCustomer(body io.ReadCloser) (customer models.Customer) {
	data , _:= ioutil.ReadAll(body);
	if err :=  json.Unmarshal(data , &customer); err != nil {
		log.Println(err);
		log.Println("there is a problem on data binding")
	}
	 return customer
}
