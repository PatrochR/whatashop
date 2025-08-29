package handler

import (
	"errors"
	"net/http"

	"github.com/PatrochR/whatashop/model/dto"
	"github.com/PatrochR/whatashop/helper"
	"github.com/PatrochR/whatashop/repository"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
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
		result.Error = errors.New("Something wrong!")
		result.Value = nil
		helper.WriteJSON(w, http.StatusInternalServerError, result)
		return
	}
	result.Value = dto.ConvertToUserGetAll(users)

	helper.WriteJSON(w, http.StatusOK, result)
}
