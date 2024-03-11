package router

import (
	"airbd/database"
	articleshandling "airbd/database/articles_handling"
	hashpassword "airbd/database/hash_password"
	"airbd/database/models"
	userhandling "airbd/database/user_handling"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user models.User

		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(jsondata, &user)
		encrypted := hashpassword.Hashthispassword(user.Password)

		user_db, err := userhandling.Finduseremail(db, user.Email)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		if encrypted == user_db.Password {
			c.JSON(http.StatusAccepted, user_db)
			return
		}
		c.String(http.StatusNotFound, "Email or password is invalid")
	}
	return fn
}

func register(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user models.User

		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(jsondata, &user)

		err = userhandling.Newuser(db, user.Name, user.Email, user.Password)
		if err != nil {
			c.String(http.StatusNotAcceptable, err.Error())
			return
		}
		c.Status(http.StatusAccepted)
	}
	return fn
}

func findname(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user models.User

		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(jsondata, &user)

		username, err := userhandling.Findusername(db, user.Id)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusFound, username)
	}
	return fn
}

func articles(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		articles := articleshandling.GetAllArticleinfo(db)

		c.JSON(http.StatusOK, articles)
	}
	return fn
}

func article(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var article models.Article
		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(jsondata, &article)
		fmt.Println(article)

		article, err = articleshandling.GetArticleinfo(db, article.Id)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusFound, article)
	}
	return fn
}

func newarticle(db *models.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var article models.Article

		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(jsondata, &article)

		err = articleshandling.Newarticle(db, article.Name, article.Size, article.Price, article.State, article.Brand, article.Color, article.UserId)
		if err != nil {
			c.String(http.StatusNotAcceptable, err.Error())
			return
		}
		c.Status(http.StatusAccepted)
	}
	return fn
}

func ApplyRoutes(r *gin.Engine) {
	db := database.Newdatabase()

	r.POST("/user_register", register(&db))
	r.POST("/user_login", login(&db))
	r.POST("/new_article", newarticle(&db))
	r.GET("/articles", articles(&db))
	r.GET("/article", article(&db))
	r.GET("/username", findname(&db))
}
