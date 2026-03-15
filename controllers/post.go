package controllers

import (
	"blogapi/config"
	"blogapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sab posts lao
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("User").Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"total": len(posts),
		"posts": posts,
	})
}

// Ek post lao
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post nahi mili"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

// Naya post banao
func CreatePost(c *gin.Context) {
	var input models.PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title aur content do",
		})
		return
	}

	// Middleware se user_id lo
	userID, _ := c.Get("user_id")

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  uint(userID.(float64)),
	}

	config.DB.Create(&post)
	c.JSON(http.StatusCreated, gin.H{
		"message": "✅ Post ban gayi!",
		"post":    post,
	})
}

// Post update karo
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post nahi mili"})
		return
	}

	// Sirf apni post update karo
	if post.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Yeh tumhari post nahi hai!",
		})
		return
	}

	var input models.PostInput
	c.ShouldBindJSON(&input)

	config.DB.Model(&post).Updates(map[string]interface{}{
		"title":   input.Title,
		"content": input.Content,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "✅ Post update ho gayi!",
		"post":    post,
	})
}

// Post delete karo
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post nahi mili"})
		return
	}

	// Sirf apni post delete karo
	if post.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Yeh tumhari post nahi hai!",
		})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "✅ Post delete ho gayi!"})
}