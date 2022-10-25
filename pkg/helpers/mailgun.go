package helpers

import (
	"fmt"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
	"os"
)

func SendActivationMail(recipient string, token string, ctx *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Err().Error(err)
		os.Exit(1)
	}

	domain := "sandbox36c8532467454f5c9326316c3e3fa652.mailgun.org"
	apiKey := "285fe120a95ca7799a1d5014126fd5a7-10eedde5-92d61a4e"

	mg := mailgun.NewMailgun(domain, apiKey)
	sender := "info@wlbtest.org"
	subject := "ACCOUNT VERIFICATION"

	message := mg.NewMessage(sender, subject, "", recipient)
	body := `
<html>
<body>
	<h1>ACCOUNT VERIFICATION</h1>
	<p style="color:blue; font-size:30px;">Hello new user</p>
	<p style="font-size:30px;">Clik this to verif your account <a href="` + pwd + `/api/auth/verification-account?verif=` + token + `">VERIF NOW</a></p>
	<p style="color:black; font-size:15px;">Or submit code ` + token + `</p>
</body>
</html>
`
	message.SetHtml(body)

	message.SetHtml(body)

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Err().Error(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func SengNotifEmail(recipient string, text string, title string, ctx *gin.Context) {
	domain := "sandbox36c8532467454f5c9326316c3e3fa652.mailgun.org"
	apiKey := "285fe120a95ca7799a1d5014126fd5a7-10eedde5-92d61a4e"

	mg := mailgun.NewMailgun(domain, apiKey)
	sender := "info@wlbtest.org"
	subject := title

	message := mg.NewMessage(sender, subject, "", recipient)
	body := `
<html>
<body>
	<h1>` + title + `</h1>
	<p style="color:black; font-size:15px;">` + text + `</p>
</body>
</html>
`
	message.SetHtml(body)

	message.SetHtml(body)

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Err().Error(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
