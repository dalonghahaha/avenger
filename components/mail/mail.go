package mail

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/spf13/viper"
)

var dialers = map[string]*gomail.Dialer{}

func Register() error {
	configs := viper.GetStringMap("component.mail")
	for key := range configs {
		config := viper.GetStringMapString("component.mail." + key)
		_port, err := strconv.Atoi(config["port"])
		if err != nil {
			return err
		}
		_tls, err := strconv.ParseBool(config["tls"])
		if err != nil {
			return err
		}
		dialer := gomail.NewDialer(config["host"], _port, config["user"], config["password"])
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: _tls}
		dialers[key] = dialer
	}
	return nil
}

func Get(key string) *gomail.Dialer {
	dialer, ok := dialers[key]
	if !ok {
		panic("db配置不存在:" + key)
	}
	return dialer
}

func BuildMessage(from string, mailTo []string, subject string, body string) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", mailTo...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	return message
}

func Send(dialer *gomail.Dialer, message *gomail.Message) error {
	return dialer.DialAndSend(message)
}
