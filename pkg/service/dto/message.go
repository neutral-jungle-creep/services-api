package dto

type MessageDto struct {
	Id       int    `json:"id"`
	SenderId int64  `json:"senderId"`
	Date     string `json:"date"`
	Text     string `json:"text"`
	IsEdit   bool   `json:"isEdit"`
}

func NewMessage(id int, senderId int64, date string, text string, isEdit bool) *MessageDto {
	return &MessageDto{
		Id:       id,
		SenderId: senderId,
		Date:     date,
		Text:     text,
		IsEdit:   isEdit,
	}
}
