package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	// "github.com/Rohanrevanth/chat-demo-go/auth"
	"github.com/Rohanrevanth/chat-demo-go/database"
	"github.com/Rohanrevanth/chat-demo-go/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON from the database.
func GetUserConversations(c *gin.Context) {
	var newConvo models.Conversation
	if err := c.BindJSON(&newConvo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conversations, err := database.GetUserConversations(newConvo.Chatuserid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, conversations)
}

// func GetAllChats(c *gin.Context) { //dev
// 	chats, err := database.GetAllChats()
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, chats)
// }

func GetUserConversation(c *gin.Context) {
	var convo models.Conversation
	if err := c.BindJSON(&convo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(convo)

	convo, err := database.GetUserConversation(convo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user conversation"})
		return
	}
	c.IndentedJSON(http.StatusOK, convo)
}

func AddNewConversation(c *gin.Context) {
	var newConvo models.Conversation
	if err := c.BindJSON(&newConvo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.AddConversation(newConvo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Added conversation successfully"})
}

func DeleteUserConversation(c *gin.Context) {
	var convo models.Conversation
	if err := c.BindJSON(&convo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	convo, err := database.DeleteUserConversation(convo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user conversation"})
		return
	}
	c.IndentedJSON(http.StatusOK, convo)
}

func AddNewMessage(c *gin.Context) {
	// Parse the multipart form (this will populate form fields and files)
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error parsing form data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form data"})
		return
	}

	// Extract the JSON message string from form values
	messageJson := form.Value["message"]
	if len(messageJson) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message data is missing"})
		return
	}

	// Parse the JSON message
	var message models.Message
	if err := json.Unmarshal([]byte(messageJson[0]), &message); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}

	// Process each file uploaded
	files := form.File["files"]
	var filePaths []string
	if files == nil {
		filePaths = models.StringArray{}
	}
	for _, file := range files {

		// Save each file (you could also process them as needed)
		if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		filePath := file.Filename
		filePaths = append(filePaths, filePath)
	}

	message.Files = filePaths
	// Process the message (e.g., save to database)
	message_, err := database.AddMessage(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a successful response
	c.JSON(http.StatusOK, message_)
}

func GetFile(c *gin.Context) {
	var fileStruct models.File
	if err := c.BindJSON(&fileStruct); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(fileStruct)

	filePath := filepath.Join("./uploads", fileStruct.Name)

	// Check if file exists and send it to the client
	if _, err := filepath.Abs(filePath); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(filePath)
}
