package config

import (
	"fmt"

	chat "github.com/davidPardoC/go-chat/internal/chat/model"
	user "github.com/davidPardoC/go-chat/internal/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConectDatabase(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Database,
		config.Database.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
}

func AutomigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &chat.Chat{}, &chat.Message{}, &chat.ChatMember{})
}
