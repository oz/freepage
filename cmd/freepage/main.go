package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/mail"
	"os"

	mp "github.com/oz/freepage"
)

var user, pass, msg string

func init() {
	flag.StringVar(&user, "u", "", "Free mobile user")
	flag.StringVar(&pass, "s", "", "Free mobile secret")
	flag.StringVar(&msg, "m", "", "Message, use \"-\" to read mails from STDIN")
}

func readMailMessage(r io.Reader) (string, error) {
	m, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		return "", err
	}
	header := m.Header
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		return "", err
	}
	txt := header.Get("Subject") + "\n" + string(body)
	return txt, nil
}

func main() {
	flag.Parse()
	if user == "" {
		log.Fatal("Set user with -u flag")
	}
	if pass == "" {
		log.Fatal("Set password with -p flag")
	}
	if msg == "" {
		log.Fatal("Set message with -m flag")
	}

	var msgText string
	if msg == "-" {
		var err error
		msgText, err = readMailMessage(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		msgText = msg
	}

	err := mp.SendSMS(user, pass, msgText)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("SMS sent")
}
