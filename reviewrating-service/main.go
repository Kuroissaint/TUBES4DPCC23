package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inisialisasi Database, Repo, dan Service
	db := ConnectDB()
	repo := NewReviewRepository(db)
	service := NewReviewService(repo)

	// 2. Siapkan Router (Jalur API) menggunakan Gin
	r := gin.Default()

	r.POST("/reviews", func(c *gin.Context) {
		var rev Review
		
		// Tangkap data JSON yang dikirim pengguna
		if err := c.ShouldBindJSON(&rev); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
			return
		}

		// Kirim ke Service untuk divalidasi dan disimpan
		err := service.SubmitReview(c.Request.Context(), rev)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Ulasan berhasil disimpan!"})
	})

	// 3. Jalankan server di port 8008 (Sesuai dengan k8s deployment)
	r.Run(":8008")
}