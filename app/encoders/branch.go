package encoders
import (
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/shireAli/GoApi/app/models"
)
func EncodeBranch(body io.ReadCloser) (branch models.Branch) {
	data , _:= ioutil.ReadAll(body);
	if err :=  json.Unmarshal(data , &branch); err != nil {
		log.Println(err);
		log.Println("there is a problem on data binding")
	}
	 return branch
}
