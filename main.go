package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/smtp"
	"os"

	"github.com/syumai/postmail/auth"
)

const (
	serverNameKey = "POSTMAIL_SMTP_SERVER_NAME"
	portKey       = "POSTMAIL_SMTP_PORT"
	userNameKey   = "POSTMAIL_SMTP_USER_NAME"
	passwordKey   = "POSTMAIL_SMTP_PASSWORD"
)

var (
	senderEmail = flag.String("f", "", "from")
)

func getAddr() (string, string, error) {
	serverName := os.Getenv(serverNameKey)
	if serverName == "" {
		return "", "", fmt.Errorf(serverNameKey + " is not set.")
	}
	port := os.Getenv(portKey)
	if port == "" {
		return "", "", fmt.Errorf(portKey + " is not set.")
	}
	return serverName, port, nil
}

func getAuth() (smtp.Auth, error) {
	userName := os.Getenv(userNameKey)
	if userName == "" {
		return nil, fmt.Errorf(userNameKey + " is not set.")
	}
	password := os.Getenv(passwordKey)
	if password == "" {
		return nil, fmt.Errorf(passwordKey + " is not set.")
	}
	return auth.LoginAuth(userName, password), nil
}

func run() error {
	serverName, port, err := getAddr()
	if err != nil {
		return err
	}
	addr := serverName + ":" + port

	authInfo, err := getAuth()
	if err != nil {
		return err
	}

	flag.Parse()
	if *senderEmail == "" {
		return fmt.Errorf("sender email address is required. provide it with `-f` flag.")
	}
	recipientEmail := flag.Arg(0)
	if recipientEmail == "" {
		return fmt.Errorf("recipient email address is required.")
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         serverName,
	}

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	c.StartTLS(tlsConfig)

	if err := c.Auth(authInfo); err != nil {
		return err
	}
	if err := c.Mail(*senderEmail); err != nil {
		return err
	}
	if err := c.Rcpt(recipientEmail); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = io.Copy(wc, os.Stdin)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
