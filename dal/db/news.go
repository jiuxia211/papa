package db

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title   string
	Content string
	Date    string
}

func CreateNews(news *News) error {
	if err := DB.Create(news).Error; err != nil {
		return err
	}
	return nil
}
