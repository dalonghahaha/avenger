package tests
import (
	"fmt"
	"testing"
	"github.com/spf13/viper"
	"avenger/components/mail"
)
func MailInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic("go fuck yourself!:" + err.Error())
	}
	err = mail.Register()
	if err != nil {
		panic("Register Fail:" + err.Error())
	}
}
func TestSendMail(t *testing.T) {
	MailInit()
	message := mail.BuildMessage("dengjialong@avenger.com", []string{"dalonghahaha@163.com"}, "测试", "测试")
	dialer := mail.Get("qq")
	err := mail.Send(dialer, message)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("发送成功")
}