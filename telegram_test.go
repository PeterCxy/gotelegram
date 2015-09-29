package telegram

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"testing"
)

func initialize() *Telegram {
	b, _ := ioutil.ReadFile("./test_key.txt")
	k := strings.Trim(string(b), " \n")
	return New(k, true)
}

func getUid() int {
	b, _ := ioutil.ReadFile("./test_user.txt")
	i, _ := strconv.ParseInt(strings.Trim(string(b), " \n"), 10, 32)
	return int(i)
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

func TestSendMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test for sending message")
	} else {
		tg := initialize()
		uid := getUid()
		fmt.Println(tg.SendMessage("tesuto", uid))
	}
}
