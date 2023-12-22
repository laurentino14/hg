package email

import (
	"github.com/resendlabs/resend-go"
	"os"
)

func GenerateResendClient() resend.EmailsSvc {
	apiKey := os.Getenv("RESEND_SECRET")

	client := resend.NewClient(apiKey)
	return client.Emails
}

type ResendEmailParams = resend.SendEmailRequest
type ResendEmailResponse = resend.SendEmailResponse
