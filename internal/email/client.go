package email

import (
	"fmt"
	"net/smtp"
)

type Client struct {
	auth smtp.Auth

	from        string
	tlsHostName string
	serverURL   string
	serverPort  string
}

func NewClient(config Config) *Client {
	return &Client{
		tlsHostName: config.TLSHostName,
		auth:        smtp.PlainAuth(config.Identity, config.Username, config.Password, config.ServerURL),
		serverURL:   config.ServerURL,
		serverPort:  config.Port,
		from:        config.From,
	}
}

func (c *Client) SendEmail(to string, message []byte) error {
	// Send without TLS
	if c.tlsHostName == "localhost" || c.tlsHostName == "" {
		err := smtp.SendMail(fmt.Sprintf("%s:%s", c.serverURL, c.serverPort), c.auth, c.from, []string{to}, message)
		if err != nil {
			return fmt.Errorf("error sending email: %w", err)
		}
		return nil
	}

	return sendMailUsingTLS(c.tlsHostName, c.serverURL, c.serverPort, c.from, to, c.auth, message)
}
