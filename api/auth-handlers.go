package api

import (
	"fmt"
	"net/http"

	login "finances.jordis.golang/application/members/auth/commands/login-user"
	register "finances.jordis.golang/application/members/auth/commands/register-user"
	verify "finances.jordis.golang/application/members/auth/commands/verify-user"

	jwtService "finances.jordis.golang/services/jwt"
	"github.com/gin-gonic/gin"
)

type ResponseCreatedUserDTO struct {
	Pk             string `json:"user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	CurrentBalance int64  `json:"current_balance"`
}

func (app *App) RegisterUserHandler(c *gin.Context) {
	var input struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Password       string `json:"password"`
		RepeatPassword string `json:"repeated_password"`
		CurrentBalance int64  `json:"current_balance"`
		MonthlyIncome  int64  `json:"monthly_income"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	command := register.RegisterUserCommand{
		Name:           input.Name,
		Mail:           input.Email,
		Password:       input.Password,
		CurrentBalance: input.CurrentBalance,
		MonthlyIncome:  input.MonthlyIncome,
	}

	user, err := register.RegisterUserCommandHandler(command, app.UsersRepo, app.Mailservice)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	responseUser := ResponseCreatedUserDTO{
		Pk:             user.Pk.Val,
		Name:           user.Username.Val,
		Email:          user.Mail.Val,
		CurrentBalance: user.CurrentBalance.Val,
	}

	jwtService := jwtService.New()
	token, err := jwtService.GenerateToken(user.Pk.Val)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Print(app.UsersRepo)

	c.JSON(http.StatusCreated, gin.H{
		"user":  responseUser,
		"token": token,
	})

}

func (app *App) LoginUserHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	command := login.LoginUserCommand{
		Mail:     input.Email,
		Password: input.Password,
	}

	token, err := login.LoginUserCommandHandler(command, app.UsersRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (app *App) VerifyUserHandler(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	command := verify.VerifyUserCommand{
		Email: input.Email,
		Code:  input.Code,
	}
	err = verify.VerifyUserCommandHandler(command, app.UsersRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"msg": "user verificated succesfully",
	})

}
