package modal

import (
	"apiProject/db"
	"apiProject/utils"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type User struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique_index:idx_email_account"`
	AccountID uint64 `json:"account_id" gorm:"unique_index:idx_email_account"`
	//Token     string     `json:"token" sql:"-"`
	//ExpiresAt time.Time  `json:"expires_at" sql:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP"`
}

type TokenClaim struct {
	UserId uint64
	jwt.StandardClaims
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) UserCreate() map[string]interface{} {

	if resp, ok := user.validate(); !ok {
		return resp
	}

	// store password as hash
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := db.GetDB().Create(user).Error

	if err != nil {
		return utils.Message(false, "Failed to create user: "+err.Error())
	}

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	response := utils.Message(true, "User has been created")
	user.Password = ""
	response["user"] = user
	return response

}

// Generate a JWT token for the user.
func LoginUser(email, password string) map[string]interface{} {

	user := &User{}
	err := db.GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Invalid login credentials. Please try again")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}

	//Worked! Logged In
	user.Password = ""

	// Declare the expiration time of the token
	// here, we have kept it as 30 minutes
	expirationTime := time.Now().Add(30 * time.Minute)

	//Create JWT token with expiry
	tokenClaim := &TokenClaim{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	// Create the JWT string with secret from config
	var byteArrayJwtKey = []byte(viper.GetString("jwt.appSecret"))
	tokenString, err := jwtToken.SignedString(byteArrayJwtKey)
	if err != nil {
		return utils.Message(false, "Error login")
	}

	//user.Token = tokenString        //Store the token in the response
	//user.ExpiresAt = expirationTime // Store expiration time

	resp := utils.Message(true, "Logged In, please use this token for auth calls")
	resp["user"] = user
	resp["token"] = tokenString
	resp["expiresAt"] = expirationTime
	return resp

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
	user.Password = ""
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

	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password is required or too short"), false
	}

	return utils.Message(false, "Requirement passed"), true
}
