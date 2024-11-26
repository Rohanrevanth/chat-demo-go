package database

import (
	"fmt"
	"log"

	"github.com/Rohanrevanth/chat-demo-go/models"
	websocket "github.com/Rohanrevanth/chat-demo-go/websockets"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "modernc.org/sqlite"
)

// var db *sql.DB

var db *gorm.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// type User struct {
// 	ID       int64  `json:"id", omitempty`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func ConnectToDB() {
	// // Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "geronimo",
	// 	// User:   os.Getenv("DBUSER"),
	// 	// Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "recordings",
	// }
	// // Get a database handle.
	// var err error
	// db, err = sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Connected!")
}

func ConnectDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Conversation{})
	db.AutoMigrate(&models.Message{})
	// db.AutoMigrate(&models.Chat{})
	fmt.Println("Connected to sqlite...")
}

// albumByID queries for the album with the specified ID.
func GetUserByEmail(email string) (models.User, error) {
	var usr models.User
	if err := db.Where("email = ?", email).First(&usr).Error; err != nil {
		return usr, fmt.Errorf("login: %v", err)
	}
	return usr, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	return users, nil
}

func GetPublicUsers() ([]models.PublicUser, error) {
	var users []models.User
	var publicUsers []models.PublicUser

	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	for _, user := range users {
		userResponse := models.PublicUser{
			Username: user.Username,
			Email:    user.Email,
			Id:       user.ID,
		}
		publicUsers = append(publicUsers, userResponse)
	}
	return publicUsers, nil
}

func SignupUser(user models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("addUser: %v", err)
	}
	return nil
}

func AddConversation(convo models.Conversation) error {
	if err := db.Create(&convo).Error; err != nil {
		return fmt.Errorf("addConvo: %v", err)
	}

	notification := models.SocketMessage{
		Type:       "FRIEND_REQUEST",
		Profileid:  convo.Chatuserid,
		Chatuserid: convo.Profileid,
	}
	websocket.SendNotificationToClients(notification)
	return nil
}

func GetUserConversations(id uint) ([]models.Conversation, error) {
	var convo []models.Conversation
	if err := db.Preload("Messages").Where("(Chatuserid = ? OR Profileid = ?) AND Isactive = ?", id, id, true).Find(&convo).Error; err != nil {
		return convo, fmt.Errorf("GetUserConversations: %v", err)
	}
	return convo, nil
}

// func GetAllChats() ([]models.Chat, error) {
// 	var chats []models.Chat
// 	if err := db.Preload("Messages").Find(&chats).Error; err != nil {
// 		return chats, fmt.Errorf("GetAllChats: %v", err)
// 	}
// 	return chats, nil
// }

// func AddChat(chat_ models.Chat) error {
// 	if err := db.Preload("Messages").Create(&chat_).Error; err != nil {
// 		return fmt.Errorf("AddChat: %v", err)
// 	}
// 	return nil
// }

func GetUserConversation(conversation models.Conversation) (models.Conversation, error) {
	var convo models.Conversation
	if err := db.Preload("Messages").Where("ID = ? AND Isactive = ?", conversation.ID, true).First(&convo).Error; err != nil {
		return convo, fmt.Errorf("GetUserConversation: %v", err)
	}

	convo.Seen = conversation.Seen

	if convo.IsFriends {

	} else {
		if conversation.IsFriends {
			convo.IsFriends = conversation.IsFriends
			notification := models.SocketMessage{
				Type:           "FRIEND_ADDED",
				Profileid:      convo.Profileid,
				Chatuserid:     convo.Chatuserid,
				COnversationID: convo.ID,
			}
			websocket.SendNotificationToClients(notification)
		}
	}

	if err := db.Save(&convo).Error; err != nil {
		return convo, fmt.Errorf("GetUserConversation: could not save updated conversation: %v", err)
	}
	return convo, nil
}

func DeleteUserConversation(conversation models.Conversation) (models.Conversation, error) {
	var convo models.Conversation
	if err := db.Preload("Messages").Where("ID = ?", conversation.ID).First(&convo).Error; err != nil {
		return convo, fmt.Errorf("DeleteUserConversation: %v", err)
	}

	convo.Isactive = false

	notification := models.SocketMessage{
		Type:           "FRIEND_REQ_REJECTED",
		Profileid:      convo.Profileid,
		Chatuserid:     convo.Chatuserid,
		COnversationID: convo.ID,
	}
	websocket.SendNotificationToClients(notification)

	if err := db.Save(&convo).Error; err != nil {
		return convo, fmt.Errorf("DeleteUserConversation: could not delete conversation: %v", err)
	}
	return convo, nil
}

func AddMessage(message models.Message) (models.Message, error) {
	var convo models.Conversation
	if err := db.Preload("Messages").Where("ID = ?", message.ChatID).First(&convo).Error; err != nil {
		return message, fmt.Errorf("AddMessage: %v", err)
	}

	convo.Messages = append(convo.Messages, message)
	convo.Seen = false

	if err := db.Save(&convo).Error; err != nil {
		return message, fmt.Errorf("AddMessage: could not save updated chat: %v", err)
	}

	var profileID uint
	if message.ProfileID == convo.Profileid {
		profileID = convo.Chatuserid
	} else {
		profileID = convo.Profileid
	}
	notification := models.SocketMessage{
		Type:           "NEW_MESSAGE",
		Profileid:      profileID,
		Chatuserid:     message.ProfileID,
		COnversationID: convo.ID,
	}
	websocket.SendNotificationToClients(notification)

	return message, nil
}

func GetFile() ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	return users, nil
}
