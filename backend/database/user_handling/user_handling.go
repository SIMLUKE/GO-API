package userhandling

import (
	hashpassword "airbd/database/hash_password"
	"airbd/database/models"
	"errors"

	"github.com/google/uuid"
)

func Findusername(db *models.Database, id string) (string, error) {
	var user models.User
	result := db.Db.First(&user, "Id = ?", id).Select(id)
	if result.RowsAffected == 0 {
		return "", errors.New("Name not found")
	}
	return user.Name, nil
}

func Finduseremail(db *models.Database, email string) (models.User, error) {
	var user models.User
	result := db.Db.First(&user, "Email = ?", email).Select(email)
	if result.RowsAffected == 0 {
		return user, errors.New("Email not found")
	}
	return user, nil
}

func Getuserinfo(db *models.Database, id string) (models.User, error) {
	var user models.User
	result := db.Db.First(&user, "Id = ?", id).Select(id)
	if result.RowsAffected == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}

func Newuser(db *models.Database, name string, email string, password string) error {
	var user models.User

	if Existmail(db, email) {
		return errors.New("Email already exists")
	}
	if name == "" || password == "" {
		return errors.New("Invalid name or password")
	}
	user.Id = uuid.NewString()
	user.Name = name
	user.Email = email
	user.Password = hashpassword.Hashthispassword(password)

	db.Db.Create(user)
	return nil
}

func Existmail(db *models.Database, email string) bool {
	result := db.Db.Where("Email = ?", email).Find(&models.User{})

	if result.RowsAffected != 0 {
		return true
	}
	return false
}
