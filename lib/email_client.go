package lib


type EmailClientInterface interface {
	SendEmail(from string, to string, subject string, body string) error
}

type EmailClient struct {}

func (emailClient *EmailClient) SendEmail(from string, to string, subject string, body string) error {
	return nil
}
