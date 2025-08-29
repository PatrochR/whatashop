package dto

import (
	"time"

	"github.com/PatrochR/whatashop/model"
)


type UserGetAll struct {
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func ConvertToUserGetAll(users *[] model.User) *[]UserGetAll{
	
	var result []UserGetAll
	for _ , e := range *users{
		var user UserGetAll
		user.Username = e.Username 
		user.Email = e.Email 
		user.CreatedAt = e.CreatedAt 
		user.UpdatedAt = e.UpdatedAt 
		result = append(result, user)
	}
	return &result

}
