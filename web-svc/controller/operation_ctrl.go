package controller

import (
	"encoding/json"
	"fmt"
	"max-web-svc/model"
	"max-web-svc/service"
	"net/http"
)

type OperationController struct {
	OperationService service.OperationService
	AuthService      service.AuthService
}

func NewOperationController(svcs *service.Services) *OperationController {
	ctrl := &OperationController{
		OperationService: svcs.OperationService,
		AuthService:      svcs.AuthService,
	}
	return ctrl
}

func (c *OperationController) Deposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var deposit model.Operation
	dErr := json.NewDecoder(r.Body).Decode(&deposit)
	if dErr != nil {
		msg := fmt.Sprint("Error [OperationController Deposit]", dErr.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if !c.AuthService.ValidateToken(ctx, deposit.Token) {
		msg := "Error [OperationController Withdraw] Token is not valid"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	balance, err := c.OperationService.Deposit(ctx, deposit)
	if err != nil {
		msg := fmt.Sprint("Error [OperationController Deposit]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(balance)
	if mErr != nil {
		msg := fmt.Sprint("Error [OperationController Deposit]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *OperationController) Withdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var withdraw model.Operation
	dErr := json.NewDecoder(r.Body).Decode(&withdraw)
	if dErr != nil {
		msg := fmt.Sprint("Error [OperationController Withdraw]", dErr.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if !c.AuthService.ValidateToken(ctx, withdraw.Token) {
		msg := "Error [OperationController Withdraw] Token is not valid"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	balance, err := c.OperationService.Withdraw(ctx, withdraw)
	if err != nil {
		msg := fmt.Sprint("Error [OperationController Withdraw]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(balance)
	if mErr != nil {
		msg := fmt.Sprint("Error [OperationController Withdraw]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
