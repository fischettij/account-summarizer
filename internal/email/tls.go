package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func sendMailUsingTLS(tlsHostName, serverURL, serverPort, from, to string, auth smtp.Auth, message []byte) error {
	tlsConfig := &tls.Config{ServerName: tlsHostName, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", serverURL, serverPort), tlsConfig)
	if err != nil {
		return fmt.Errorf("error connecting to SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, tlsHostName)
	if err != nil {
		return fmt.Errorf("error creating to SMTP server: %w", err)
	}
	defer client.Close()

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("error authentication: %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("error setting mail sender: %w", err)
	}

	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("error setting mail recipient: %w", err)
	}

	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("error in message open: %w", err)
	}
	defer wc.Close()

	_, err = wc.Write(message)
	//_, err = fmt.Fprintf(wc, string(message))
	if err != nil {
		return fmt.Errorf("error writing message: %w", err)
	}

	return nil
}
