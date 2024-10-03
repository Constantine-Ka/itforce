package api

import (
	"ITForce/internal/repositories"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

type ErrorMsg struct {
	Message   string  `json:"message"`
	Mode      string  `json:"mode"`
	UserId    string  `json:"userId"`
	Amount    string  `json:"amount"`
	Balance   float64 `json:"balance"`
	HistoryId int64   `json:"historyId"`
}
type ErrorShort struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
}

var lock = sync.Mutex{}

// PurchaseHandler
//
//	@ID	 purchase-handler
//	@Summary Изменяет баланс. Использует Мьютексы, для блокировки параллельного изменения данных
//	@Produce json
//	@Param user query     string true "Идентификатор пользователя"
//	@Param amount query     string true "Сумма списания"
//
// @Success 200 {object} api.ErrorMsg
// @Failure 400 {object} api.ErrorMsg
// @Router /purchase [post]
func PurchaseHandler(c echo.Context, db *sqlx.DB) error {
	user := c.FormValue("user")
	amountStr := c.FormValue("amount")
	errMSG := ErrorMsg{
		Amount: amountStr,
		UserId: user,
	}
	if user == "" || amountStr == "" {
		c.Logger().Error("badRequest")
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	userId, err := strconv.Atoi(user)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if amount < 0 {
		amount *= -1
	}
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	lock.Lock()
	defer lock.Unlock()
	userBalance, err := repositories.GetBalanceById(c.Request().Context(), db, userId)
	if err != nil {
		errMSG.Message = err.Error()
		errMSG.Mode = "Ошибка SQL запроса"
		c.Logger().Error(errMSG)
		return c.JSON(http.StatusOK, errMSG)
	}
	if userBalance < amount {
		errMSG.Message = "Баланс меньше суммы платежа"
		errMSG.Mode = "Логическая ошибка"
		c.Logger().Error(errMSG)
		return c.JSON(http.StatusOK, errMSG)
	}
	newBalance, err := repositories.UpdateAmount(c.Request().Context(), db, userId, amount)
	if err != nil {
		errMSG.Message = err.Error()
		c.Logger().Error(errMSG)
		return c.JSON(http.StatusOK, errMSG)
	}
	errMSG.HistoryId, err = repositories.InsertHistory(c.Request().Context(), db, userId, 1, amount)
	if err != nil {
		errMSG.Message = err.Error()
		errMSG.Mode = "Ошибка SQL запроса"
		c.Logger().Error(errMSG)
		return c.JSON(http.StatusOK, errMSG)
	}
	errMSG.Message = fmt.Sprintf("Ваш новый Баланс %v $", newBalance)
	errMSG.Mode = "Success"
	errMSG.Balance = newBalance
	c.Logger().Error(errMSG)
	return c.JSON(http.StatusOK, errMSG)
}

// HistoryHandler
//
//		@ID	 pistory-handler
//		@Summary Показывает историю по одному или всем пользователям
//		@Produce json
//		@Param  id query string true "Идентификатор пользователя"
//	 @Success 200 {array} repositories.ActionHistory
//	 @Failure 400 {object} api.ErrorShort
//	 @Router /history [get]
func HistoryHandler(c echo.Context, db *sqlx.DB) error {
	idStr := c.Param("id")
	if idStr == "" {
		idStr = c.QueryParam("id")
	}
	errMsg := ErrorShort{UserId: idStr}
	id, _ := strconv.Atoi(idStr)

	history, err := repositories.GetHistory(c.Request().Context(), db, id)
	if err != nil {
		errMsg.Message = err.Error()
		c.Logger().Error(errMsg)
		return c.JSON(http.StatusOK, errMsg)
	}
	return c.JSON(http.StatusOK, history)
}
