package api

import (
	"net/http"

	groups_useCases "finances.jordis.golang/application/members/groups/create-group"
	"github.com/gin-gonic/gin"
)

type GroupCreatedResponse struct {
	Pk          string `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Secret      string `json:"secret"`
}

func (app *App) CreateGroup(c *gin.Context) {

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no token found",
		})
		return
	}

	command := groups_useCases.CreateGroupCommand{
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   input.UserID,
	}

	group, err := groups_useCases.CreateGroupcommandHandler(command, app.GroupsRepo, app.UsersGroupsRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no token found",
		})
		return
	}

	groupRespnse := GroupCreatedResponse{
		Pk:          group.Pk.Val,
		Name:        group.Name.Val,
		Description: group.Description.Val,
		Secret:      group.Secret.Val,
	}

	c.JSON(http.StatusCreated, gin.H{
		"group created": groupRespnse,
	})

}
