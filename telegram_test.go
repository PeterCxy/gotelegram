package telegram

import (
	"io/ioutil"
	"strings"
	"fmt"
	"testing"
)

func initialize() *Telegram {
	b, _ := ioutil.ReadFile("./test_key.txt")
	k := strings.Trim(string(b), " \n")
	return New(k, true)
}

func TestSetWebhook(*testing.T) {
	tg := initialize()

	fmt.Println(tg.SetWebhook("https://google.com/"))
	fmt.Println(tg.SetWebhook(""))
}

func TestGetUpdates(*testing.T) {
	tg := initialize()

	fmt.Println(tg.GetUpdates(0, 100, 5))
}
