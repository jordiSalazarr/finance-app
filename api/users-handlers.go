package api

import (
	"fmt"
	"net/http"

	getuser "finances.jordis.golang/application/members/users/queries/get-user"
	"github.com/gin-gonic/gin"
)

// DTO for the user
type UserDTO struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	Mail           string `json:"mail"`
	CurrentBalance int64  `json:"current_balance"`
	MonthlyIncome  int64  `json:"monthly_income"`
	IsActive       bool   `json:"is_active"`
	IsVerified     bool   `json:"is_verified"`
}

func (app *App) TellMyName(c *gin.Context) {
	val, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "i dont know",
		})
		return
	}
	userId := val.(string)
	app.Logger.Info("user id", userId)

	query := getuser.GetUserQuery{
		UserID: userId,
	}

	user, err := app.UsersRepo.GetById(query.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": fmt.Sprintf("error getting user: %v", err),
		})
		return
	}
	userDTO := UserDTO{
		ID:             user.Pk.Val,
		UserName:       user.Username.Val,
		Mail:           user.Mail.Val,
		CurrentBalance: user.CurrentBalance.Val,
		MonthlyIncome:  user.MonthlyIncome.Val,
		IsActive:       user.IsActive,
		IsVerified:     user.IsVerified,
	}

	c.JSON(http.StatusOK, gin.H{
		"user": userDTO,
	})

}
