package mail_service

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

type SMPTService struct {
	Host     string
	Port     int
	User     string
	Password string
	Mailer   *mail.Dialer
	// BaseURL string // por ejemplo: https://tusitio.com/verificar

}

func New() (*SMPTService, error) {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	dialer := mail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	smptService := SMPTService{
		Host:     dialer.Host,
		Port:     port,
		User:     dialer.Host,
		Password: dialer.Password,
		Mailer:   dialer,
	}

	return &smptService, nil
}
func (smpt *SMPTService) DialAndSend(msg *mail.Message) error {
	return smpt.Mailer.DialAndSend(msg)
}

func (smpt *SMPTService) SendVerificationCode(recipient string, code string) {
	// Create a new message
	message := mail.NewMessage()

	// Set email headers
	message.SetHeader("From", "financeApp@company.com")
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", "🔐 Verifica tu cuenta")
	htmlBody := fmt.Sprintf(`
		<html>
		<body style="font-family: sans-serif; background-color: #f4f4f4; padding: 30px;">
			<div style="max-width: 600px; margin: auto; background: white; border-radius: 10px; padding: 20px; box-shadow: 0 4px 12px rgba(0,0,0,0.1);">
				<h2 style="color: #333;">¡Bienvenido/a a nuestra app de finanzas! 🎉</h2>
				<p>Tu código de verificación es:</p>
				<p style="font-size: 24px; font-weight: bold; color: #4CAF50;">%s</p>
				<p>Puedes introducirlo directamente o hacer clic en el siguiente botón:</p>
				<a href="%s/%s" style="display: inline-block; background-color: #4CAF50; color: white; padding: 10px 20px; border-radius: 5px; text-decoration: none; margin-top: 10px;">Verificar cuenta</a>
				<p style="margin-top: 30px; font-size: 12px; color: #777;">Si no has solicitado este correo, simplemente ignóralo.</p>
			</div>
		</body>
		</html>
	`, code, "www.example.com", recipient)

	message.SetBody("text/html", htmlBody)

	// Send email
	if err := smpt.DialAndSend(message); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent to:", recipient)
	}
}
