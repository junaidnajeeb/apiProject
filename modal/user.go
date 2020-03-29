package modal

import (
	"apiProject/db"
	"apiProject/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Email     string     `json:"email" gorm:"unique_index:idx_email_account"`
	AccountID uint64     `json:"account_id" gorm:"unique_index:idx_email_account"`
	Token     string     `json:"token";sql:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
}

type Token struct {
	UserId uint64
	jwt.StandardClaims
}

func (user *User) UserCreate() map[string]interface{} {

	if resp, ok := user.validate(); !ok {
		return resp
	}

	err := db.GetDB().Create(user).Error

	if err != nil {
		return utils.Message(false, "Failed to create user: "+err.Error())
	}

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	response := utils.Message(true, "User has been created")
	response["user"] = user
	return response

}

func (user *User) GetOneUser(id uint64) map[string]interface{} {

	err := db.GetDB().First(&user, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "User not found")
		}
		return utils.Message(false, "Connection error. "+err.Error())
	}

	response := utils.Message(true, "Success")
	response["user"] = user

	return response

}

func (user *User) DeleteOneUser(id uint64) map[string]interface{} {

	err := db.GetDB().First(&user, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "User not found")
		}
		return utils.Message(false, "Connection error. "+err.Error())
	}

	db.GetDB().Delete(&user)
	response := utils.Message(true, "Success")
	response["user"] = user

	return response

}

func (user *User) validate() (map[string]interface{}, bool) {

	if len(user.Name) < 1 {
		return utils.Message(false, "User name is required"), false
	}

	return utils.Message(false, "Requirement passed"), true
}
