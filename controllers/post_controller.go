package controllers

import (
	"cms-backend/models"
	"cms-backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPosts retrieves all posts with optional filtering
func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post

	title := c.Query("title")
	author := c.Query("author")

	query := db
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}

	// Use proper preloading for media relationships
	if err := query.Preload("Media").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost retrieves a specific post by ID
func GetPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	idString := c.Param("id")

	postId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "id must be an integer",
		})
		return
	}

	var post models.Post
	if err := db.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("post with id %d not found", postId),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := post.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	tx := db.Begin()
	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	tx.Commit()

	c.JSON(http.StatusCreated, post)
}

// UpdatePost updates an existing post
func UpdatePost(c *gin.Context) {
	// TODO: Get database connection
	// Get database instance from Gin context

	// TODO: Parse and validate ID
	// Get ID from URL parameter
	// Example: /api/posts/123

	// TODO: Check if post exists
	// Find existing post

	// TODO: Parse and validate update data
	// Define variable for update input

	// TODO: Update post fields
	// Update only the fields that are allowed to be updated

	// TODO: Save updates to database
	// 1. Start transaction

	// 2. Save the updated post

	// 3. Commit transaction

	// 4. Return updated post
}

// DeletePost deletes a post
func DeletePost(c *gin.Context) {
	// TODO: Get database connection
	// Get database instance from Gin context

	// TODO: Parse and validate ID
	// Get ID from URL parameter
	// Example: /api/posts/123

	// TODO: Check if post exists
	// Find existing post

	// TODO: Delete post from database
	// 1. Start transaction

	// 2. Delete the post
	// Note: Consider if you want soft delete (recommended) or hard delete

	// 3. Commit transaction

	// 4. Return success message
}
