package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mytodolist1/be_p3/model"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var res Response

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		res = SendMessage(user)
	} else {
		res = Response{
			Status:  false,
			Message: "Secret tidak valid",
		}
	}

	json.NewEncoder(w).Encode(res)

	fmt.Fprintf(w, res.Message, res.Status)
	return
}
