package db

import (
	"github.com/colebowl/coles-stream/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dsn := "host=localhost user=youruser password=yourpassword dbname=coles_stream port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate the schema
	return DB.AutoMigrate(&models.Post{}, &models.Thought{}, &models.Tag{})
}

func CreatePost(post *models.Post) error {
	return DB.Create(post).Error
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := DB.Preload("Thoughts").Preload("Tags").First(&post, id).Error
	return &post, err
}

func GetLatestPosts(limit int) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Preload("Thoughts").Preload("Tags").Order("created_at desc").Limit(limit).Find(&posts).Error
	return posts, err
}

// Add more database operations as needed
