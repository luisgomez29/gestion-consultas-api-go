package mailers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/mail"
	"net/smtp"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
)

type (
	// Template especifica el contexto y la ruta de la plantilla HTML. Por defecto se busca las plantillas
	// en la carpeta "templates/emails/".
	Template struct {
		Name    string
		Context map[string]interface{}
	}

	EmailMessage struct {
		To       mail.Address
		Subject  string
		Template Template
	}
)

// smtpConfig define el servidor y el acceso para enviar el email.
type smtpConfig struct {
	host     string
	port     string
	username string
	password string
}

// Address retorna el host y el puerto
func (s *smtpConfig) Address() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}

func Send(em *EmailMessage) (bool, error) {
	from := mail.Address{
		Name:    config.Load("DEFAULT_FROM_EMAIL"),
		Address: config.Load("EMAIL_HOST_USER"),
	}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = em.To.String()
	headers["Subject"] = em.Subject
	headers["Content-Type"] = `text/html; charset="UTF8"`

	// Message
	message := ""
	for i, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", i, v)
	}

	t, err := template.ParseFiles(fmt.Sprintf("templates/emails/%s", em.Template.Name))
	if err != nil {
		return false, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, &em.Template.Context); err != nil {
		return false, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	message += buf.String()

	// smtp configuration
	smtpCon := smtpConfig{
		host:     config.Load("EMAIL_HOST"),
		port:     config.Load("EMAIL_PORT"),
		username: config.Load("EMAIL_HOST_USER"),
		password: config.Load("EMAIL_HOST_PASSWORD"),
	}

	// Authentication
	auth := smtp.PlainAuth("", smtpCon.username, smtpCon.password, smtpCon.host)

	// Sending email
	if err = smtp.SendMail(smtpCon.Address(), auth, from.Address, []string{em.To.Address}, []byte(message)); err != nil {
		return false, err
	}
	return true, nil
}
