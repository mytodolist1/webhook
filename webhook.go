package webhook

import (
	"os"

	"github.com/aiteung/atapi"
	model "github.com/mytodolist1/be_p3/model"
)

func SendMessage(user model.User) (res Response) {
	req := Request{
		To:      user.Email,
		Message: "Hello " + user.Username + " kamu sudah berhasil register di MyTodoList, silahkan login atau tekan link ini https://mytodolist.my.id/login.html untuk melanjutkan.",
	}

	res, _ = atapi.PostStructWithToken[Response]("Token", os.Getenv("TOKEN"), req, "https://api.wa.my.id/api/send/message/text")

	return res
}
