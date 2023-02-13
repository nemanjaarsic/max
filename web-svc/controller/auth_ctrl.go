package controller

import (
	"encoding/json"
	"fmt"
	"max-web-svc/model"
	"max-web-svc/service"
	"net/http"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(svcs *service.Services) *AuthController {
	ctrl := &AuthController{
		AuthService: svcs.AuthService,
	}
	return ctrl
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var login model.Login

	dErr := json.NewDecoder(r.Body).Decode(&login)
	if dErr != nil {
		msg := fmt.Sprint("Error [AuthController Login]", dErr.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	id, err := c.AuthService.Login(ctx, login)
	if err != nil {
		msg := fmt.Sprint("Error [AuthController Login]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	token, err := c.AuthService.GetToken(ctx, id)

	response, mErr := json.Marshal(token)
	if mErr != nil {
		msg := fmt.Sprint("Error [AuthController Login]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
