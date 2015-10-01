// Convenient functions
package telegram

type TObject map[string]interface{}

func (this TObject) UpdateId() int64 {
	return int64(this["update_id"].(float64))
}

func (this TObject) Message() TObject {
	return TObject(this["message"].(map[string]interface{}))
}

func (this TObject) MessageId() int64 {
	return int64(this["message_id"].(float64))
}

func (this TObject) ChatId() int64 {
	return int64(this["chat"].(map[string]interface{})["id"].(float64))
}

func (this TObject) FromId() int64 {
	return int64(this["from"].(map[string]interface{})["id"].(float64))
}

func (this TObject) From() TObject {
	return TObject(this["from"].(map[string]interface{}))
}
