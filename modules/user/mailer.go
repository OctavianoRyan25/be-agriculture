package user

import "os"

var (
	//Testing
	EMAIL_FROM    = os.Getenv("EMAIL_FROM")
	SMTP_HOST     = os.Getenv("SMTP_HOST")
	SMTP_PORT     = os.Getenv("SMTP_PORT")
	SMTP_USER     = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASS")
)
