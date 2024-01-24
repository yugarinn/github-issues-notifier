package internal

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func sendIssueEmailAlertTo(listener Listener, issue GithubIssue) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpFrom := os.Getenv("SMTP_EMAIL_FROM")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpAuth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)

	headers := map[string]string{
		"From": smtpFrom,
		"To": listener.EmailTo,
		"Subject": fmt.Sprintf("New Issue in %s: %s", listener.Name, issue.Title),
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	header := ""
	for k, v := range headers {
		header += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body := header + "\r\n" + buildAlertEmailBody(listener, issue)

	err := smtp.SendMail(smtpHost + ":" + smtpPort, smtpAuth, smtpFrom, []string{listener.EmailTo}, []byte(body))
	if err != nil {
		log.Fatal("Error sending email: ", err)
		return err
	}

	return nil
}

func buildAlertEmailBody(listener Listener, issue GithubIssue) string {
	return fmt.Sprintf("New issue in %s<br/><br/><a href=\"%s\" target=\"_blank\">%s</a>", listener.Name, issue.Url, issue.Title)
}
