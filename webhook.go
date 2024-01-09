package webhook

import (
	"os"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	user "github.com/mytodolist1/be_p3/model"
	"github.com/whatsauth/wa"
)

func SendMessage(msg model.IteungMessage) (resp atmessage.Response) {

	var user user.User

	msg.Phone_number = user.Phonenumber

	phone := msg.Phone_number

	dt := &wa.TextMessage{
		To:       phone,
		IsGroup:  false,
		Messages: "Hello " + user.Username + " kamu sudah berhasil register di MyTodoList, silahkan login atau tekan link ini https://mytodolist.my.id/login.html untuk melanjutkan.",
	}

	resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

	return
}
