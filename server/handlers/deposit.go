package handlers

import (
	"errors"
	"strconv"

	"github.com/GOsling-Inc/GOsling/models"
	"github.com/GOsling-Inc/GOsling/services"
	"github.com/labstack/echo/v4"
)

type IDepositHandler interface {
	POST_NewDeposit(echo.Context) error
	GET_User_Deposits(echo.Context) error
}

type DepositHandler struct {
	service *services.Service
}

func NewDepositHandler(s *services.Service) *DepositHandler {
	return &DepositHandler{
		service: s,
	}
}

func (h *DepositHandler) POST_NewDeposit(c echo.Context) error {
	beta_depos := &models.Deposit{
		AccountId: c.FormValue("AccountId"),
		Period:    c.FormValue("Period"),
	}
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	_, err = h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	acc, err := h.service.GetAccountById(beta_depos.AccountId)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	beta_depos.Amount, _ = strconv.ParseFloat(c.FormValue("Amount"), 64)
	beta_depos.Percent, _ = strconv.ParseFloat(c.FormValue("Percent"), 64)
	if acc.UserId != id || acc.Amount < beta_depos.Amount {
		return c.JSON(401, errors.New("account error"))
	}
	beta_depos.UserId = id
	if err = h.service.CreateDeposit(beta_depos); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *DepositHandler) GET_USER_DEPOSITS(c echo.Context) error {
	header := c.Request().Header
	id, err := h.service.ParseJWT(header["Token"][0])
	if err != nil {
		return c.JSON(401, err.Error())
	}
	_, err = h.service.GetUser(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	depos, err := h.service.GetUserDeposits(id)
	if err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(200, depos)
}
