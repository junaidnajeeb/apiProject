package modal

import (
	"apiProject/db"
	"apiProject/utils"
	"time"

	"github.com/jinzhu/gorm"
)

const statusActive = "active"

//a struct to rep account
// id, created_at, deleted_at, updated_at is created by gorm
type Account struct {
	//gorm.Model
	ID          uint       `json:"id" gorm:"primary_key"`
	Name        string     `json:"name" gorm:"NOT NULL"`
	Description string     `json:"description" gorm:"type:TEXT"`
	Status      string     `json:"status" gorm:"type:ENUM('active', 'inactive', 'deleted') NOT NULL DEFAULT 'active'"`
	Users       []User     `json:"users" gorm:"foreignkey:AccountID"`
	CreatedAt   time.Time  `json:"created_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
}

func (account *Account) AccountCreate() map[string]interface{} {

	if resp, ok := account.validate(); !ok {
		return resp
	}

	err := db.GetDB().Create(account).Error

	if err != nil {
		return utils.Message(false, "Failed to create account: "+err.Error())
	}

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	response := utils.Message(true, "Account has been created")
	response["account"] = account
	return response

}

func (account *Account) GetOneAccount(id uint64) map[string]interface{} {

	err := db.GetDB().First(&account, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Account not found")
		}
		return utils.Message(false, "Connection error. "+err.Error())
	}

	response := utils.Message(true, "Success")
	response["account"] = account

	return response

}

func (account *Account) DeleteOneAccount(id uint64) map[string]interface{} {

	err := db.GetDB().First(&account, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Account not found")
		}
		return utils.Message(false, "Connection error. "+err.Error())
	}

	db.GetDB().Delete(&account)
	response := utils.Message(true, "Success")
	response["account"] = account

	return response

}

func (account *Account) validate() (map[string]interface{}, bool) {

	if len(account.Name) < 1 {
		return utils.Message(false, "Name is required"), false
	}

	if len(account.Status) == 0 {
		account.Status = statusActive
	}

	return utils.Message(false, "Requirement passed"), true
}
