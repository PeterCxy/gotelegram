// Telegram Bot API client implementation
package telegram

import (
	"fmt"
	"encoding/json"

	"github.com/ddliu/go-httpclient"
)

// https://api.telegram.org/bot<token>/METHOD_NAME
const MethodUrl = "https://api.telegram.org/bot%s/%s"

type Telegram struct {
	apiKey string
	debug bool
}

func New(key string, debug bool) *Telegram {
	return &Telegram {
		apiKey: key,
		debug: debug,
	}
}

func (tg *Telegram) post(method string, params map[string]string) map[string]interface{} {
	url := fmt.Sprintf(MethodUrl, tg.apiKey, method)

	if tg.debug {
		fmt.Println(url)
	}

	res, err := httpclient.Post(url, params)

	if (err != nil) || (res == nil) {
		return nil
	} else {
		var ret interface{}
		var str string
		str, err = res.ToString()

		if tg.debug {
			fmt.Println(str)
		}

		if err != nil {
			return nil
		}

		err = json.Unmarshal([]byte(str), &ret)

		if (err == nil) && (ret != nil) {
			return ret.(map[string]interface{})
		} else {
			return nil
		}
	}
}

func (tg *Telegram) SetWebhook(url string) bool {
	res := tg.post("setWebhook", map[string]string {
		"url": url,
	})

	if res == nil {
		return false
	}

	return res["ok"].(bool)
}

// timeout: time to wait (seconds)
func (tg *Telegram) GetUpdates(offset int64, limit int, timeout int) []interface{} {
	res := tg.post("getUpdates", map[string]string {
		"offset": fmt.Sprintf("%d", offset),
		"limit": fmt.Sprintf("%d", limit),
		"timeout": fmt.Sprintf("%d", timeout),
	})

	if (res == nil) || (!res["ok"].(bool)) {
		return nil
	}

	return res["result"].([]interface{})
}

// The Message object has too many fields, let's just use a map
// Consult <https://core.telegram.org/bots/api#sendmessage> for a full list of params
func (tg *Telegram) SendMessageRaw(msg map[string]string) bool {
	res := tg.post("sendMessage", msg)

	if res == nil {
		return false
	}

	return res["ok"].(bool)
}

func (tg *Telegram) SendMessage(text string, chat int64) bool {
	return tg.SendMessageRaw(map[string]string {
		"chat_id": fmt.Sprintf("%d", chat),
		"text": text,
	})
}

func (tg *Telegram) ReplyToMessage(id int64, text string, chat int64) bool {
	return tg.SendMessageRaw(map[string]string {
		"chat_id": fmt.Sprintf("%d", chat),
		"text": text,
		"reply_to_message_id": fmt.Sprintf("%d", id),
	})
}

// See <https://core.telegram.org/bots/api#sendchataction> for a list of actions
func (tg *Telegram) SendChatAction(action string, chat int64) bool {
	res := tg.post("sendChatAction", map[string]string {
		"chat_id": fmt.Sprintf("%d", chat),
		"action": action,
	})

	if res == nil {
		return false
	}

	return res["ok"].(bool)
}
