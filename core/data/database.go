package data

type UserData struct {
	ID       string        `bson:"id"`
	Messages []UserMessage `bson:"messages"`
}

func (u UserData) SetId(id string) UserData {
	u.ID = id
	return u
}

func (u UserData) AddMessage(message UserMessage) UserData {
	u.Messages = append(u.Messages, message)
	return u
}

type UserMessage struct {
	Content   string `bson:"content"`
	Timestamp int64  `bson:"timestamp"`
	Guild     string `bson:"guild"`
	Channel   string `bson:"channel"`
	ID        string `bson:"id"`
}

func NewUserMessage(content, guild, channel, id string, timestamp int64) UserMessage {
	return UserMessage{content, timestamp, guild, channel, id}
}
