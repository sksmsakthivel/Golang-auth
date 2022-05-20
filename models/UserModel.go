package models

import (
	"fmt"
	"goauth/config"
	"net/url"
	"time"
	// "github.com/flamego/validator"
)

type User struct {
	Id          int
	Name        string    `json:"name" gorm:"default:null"`
	UserName    string    `json:"userName" gorm:"default:null"`
	Password    string    `json:"-" gorm:"default:null"`
	DOB         time.Time `json:"-" gorm:"default:null"`
	DateOfBirth string    `json:"dob" gorm:"-"`
	LinkedinUrl string    `json:"linkedinUrl" gorm:"default:null"`
}

var DB = config.SetUpDB()

func Validate(u User) []string {
	fmt.Println("user:", u)
	var msg = []string{}
	if u.Name == "" {
		msg = append(msg, "Name is Required")
	}
	if u.UserName == "" {
		msg = append(msg, "User Name is Required")
	}

	if u.DateOfBirth == "" {
		msg = append(msg, "DOB is Required")
	} else {
		dob, _ := time.Parse("02-01-2006", u.DateOfBirth)
		if dob.IsZero() {
			msg = append(msg, "DOB is invalid")
		}
	}
	if u.LinkedinUrl == "" {
		msg = append(msg, "LinkedIn Url is Required")
	} else {
		_, err := url.ParseRequestURI(u.LinkedinUrl)
		if err != nil {
			msg = append(msg, "Invalid Url")
		}
	}
	return msg

}

/*create user*/
func CreateUser(user *User) (successUser User, err error) {
	if err = DB.Table("users").Create(user).Order("id desc").Take(&successUser).Error; err != nil {
		return User{}, err
	}
	return successUser, nil
}

/*find user by where condition*/
func FindUser(findOption map[string]interface{}) (user User, err error) {
	query := DB.Table("users")
	for key, val := range findOption {
		query = query.Where(key+"=?", val)
	}
	if err = query.Take(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
