package freepage

import (
	"errors"
	"net/http"
	"net/url"
)

const APIURL = "https://smsapi.free-mobile.fr/sendmsg"

func SendSMS(user, pass, txt string) error {
	msg := url.QueryEscape(txt)
	url := APIURL + "?user=" + user + "&pass=" + pass + "&msg=" + msg
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return errors.New("Missing parameter")
	case http.StatusPaymentRequired:
		return errors.New("Too many messages sent, wait a moment")
	case http.StatusForbidden:
		return errors.New("Forbidden")
	case http.StatusInternalServerError:
		return errors.New("Server error")
	}
	return nil
}
