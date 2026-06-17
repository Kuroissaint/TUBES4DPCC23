package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := ConnectDB()
	repo := NewChatRepository(db)
	service := NewChatService(repo)

	r := gin.Default()

	r.POST("/chat/send", func(c *gin.Context) {
		var msg Message
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
			return
		}

		err := service.SendMessage(c.Request.Context(), msg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pesan terkirim!"})
	})

	// Gunakan port 8009 agar tidak bentrok
	r.Run(":8009")
}