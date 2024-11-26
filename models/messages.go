package models

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Custom type for PostgreSQL array of strings (TEXT[])
type StringArray []string

type Message struct {
	gorm.Model
	Message   string      `json:"message"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	ProfileID uint        `json:"profileid"`
	Timestamp string      `json:"timestamp"`
	ChatID    uint        `json:"chatid"`
	Files     StringArray `json:"files" gorm:"type:text[]"`
}

type Conversation struct {
	gorm.Model
	Username      string `json:"username"`
	Email         string `json:"email"`
	Profileid     uint   `json:"profileid"`
	Chatusername  string `json:"chatusername"`
	Chatuseremail string `json:"chatuseremail"`
	Chatuserid    uint   `json:"chatuserid"`
	Seen          bool   `json:"seen"`
	IsFriends     bool   `json:"isfriends"`
	// LastMessage string    `json:"lastmessage"`
	// Timestamp   string    `json:"timestamp"`
	Messages []Message `json:"messages" gorm:"foreignKey:ChatID"`
	Isactive bool      `json:"isactive" gorm:"default:true"`
}

type SocketMessage struct {
	gorm.Model
	Type           string `json:"type"`
	Data           string `json:"message"`
	Profileid      uint   `json:"profileid"`
	Chatuserid     uint   `json:"chatuserid"`
	COnversationID uint   `json:"conversationid"`
}

type File struct {
	gorm.Model
	Name           string `json:"name"`
	Messageid      uint   `json:"messageid"`
	COnversationID uint   `json:"conversationid"`
}

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}
	switch v := value.(type) {
	case string:
		// If the value is an empty string, set it to an empty slice
		if v == "" {
			*s = StringArray{}
		} else {
			*s = StringArray(strings.Split(v, ","))
		}
	case []byte:
		strValue := string(v)
		if strValue == "" {
			*s = StringArray{}
		} else {
			*s = StringArray(strings.Split(strValue, ","))
		}
	default:
		return fmt.Errorf("unsupported value type %T for StringArray", v)
	}
	return nil
}

func (s StringArray) Value() (driver.Value, error) {
	// Convert the slice of strings to a single string for storage in the database
	return strings.Join(s, ","), nil
}
