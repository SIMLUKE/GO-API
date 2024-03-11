package articleshandling

import (
	"airbd/database/models"
	userhandling "airbd/database/user_handling"
	"errors"

	"github.com/google/uuid"
)

func Newarticle(db *models.Database, name string, size string, price float32, state string, brand string, color string, seller_id string) error {
	var article models.Article

	if name == "" || size == "" || state == "" || brand == "" || color == "" || seller_id == "" {
		return errors.New("Invalid article")
	}
	_, err := userhandling.Findusername(db, seller_id)
	if err != nil {
		return err
	}
	article.Name = name
	article.State = state
	article.Size = size
	article.Brand = brand
	article.Color = color
	article.Price = price
	article.UserId = seller_id
	article.Id = uuid.NewString()

	db.Db.Create(article)
	return nil
}

func GetAllArticleinfo(db *models.Database) []models.Article {
	var articles []models.Article

	db.Db.Find(&articles)

	return articles
}

func GetArticleinfo(db *models.Database, id string) (models.Article, error) {
	var article models.Article

	result := db.Db.First(&article, "Id = ?", id).Select(id)
	if result.RowsAffected == 0 {
		return article, errors.New("Article not found")
	}
	return article, nil
}
