package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var res atmessage.Response

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println(err)
	}

	if os.Getenv("SECRET") != "" {
		res = SendMessage(msg)
	} else {
		res.Response = "Secret not match"
	}

	json.NewEncoder(w).Encode(res)

	fmt.Fprintf(w, res.Response)
	return
}
