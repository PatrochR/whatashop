package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PatrochR/whatashop/helper"
	"github.com/PatrochR/whatashop/model"
	"github.com/PatrochR/whatashop/model/dto"
	"github.com/PatrochR/whatashop/repository"
	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	logger   *log.Logger
	userRepo repository.UserRepository
}

func NewUserHandler(logger *log.Logger, userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result := helper.Result{
		IsSuccess: true,
		Error:     nil,
	}
	users, err := h.userRepo.GetAll()
	if err != nil {
		result.IsSuccess = false
		result.Error = &helper.ErrSomthingWrong
		result.Value = nil
		helper.WriteJSON(w, http.StatusInternalServerError, result)
		result.Log(h.logger)
		return
	}
	result.Value = dto.ConvertToUserGetAll(users)

	helper.WriteJSON(w, http.StatusOK, result)
}

func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	result := helper.Result{
		IsSuccess: true,
		Error:     nil,
		Value:     nil,
	}
	var dto dto.UserAdd
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		result.IsSuccess = false
		result.Error = &helper.ErrSomthingWrong
		helper.WriteJSON(w, http.StatusInternalServerError, result)
		result.Log(h.logger)
		return
	}
	var user model.User
	user.Username = dto.Username
	user.Email = dto.Email
	hashed , err:= bcrypt.GenerateFromPassword([]byte(dto.Password) , 10)
	if err != nil{
		result.IsSuccess = false
		result.Error = &helper.ErrSomthingWrong
		helper.WriteJSON(w, http.StatusInternalServerError, result)
		result.Log(h.logger)
		return 
	}
	user.PasswordHash = string(hashed)
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	err = h.userRepo.Add(&user)
	if err != nil {
		result.IsSuccess = false
		result.Error = &helper.ErrSomthingWrong
		helper.WriteJSON(w, http.StatusInternalServerError, result)
		result.Log(h.logger)
		return
	}

	result.Value = "User Added"
	helper.WriteJSON(w, http.StatusOK, result)
}
