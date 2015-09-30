// Convenient functions
package telegram

type TObject map[string]interface{}

func (this TObject) MessageId()int64 {
	return int64(this["message_id"].(float64))
}

func (this TObject) ChatId() int64 {
	return int64(this["chat"].(map[string]interface{})["id"].(float64))
}
