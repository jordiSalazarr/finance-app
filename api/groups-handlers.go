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
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	var input struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no token found",
		})
		return
	}

	command := groups_useCases.CreateGroupCommand{
		Name:      input.Name,
		CreatedBy: userId,
	}

	group, err := groups_useCases.CreateGroupcommandHandler(command, app.GroupsRepo, app.UsersGroupsRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no token found",
		})
		return
	}

	groupRespnse := GroupCreatedResponse{
		Pk:     group.Pk.Val,
		Name:   group.Name.Val,
		Secret: group.Secret.Val,
	}

	c.JSON(http.StatusCreated, gin.H{
		"group created": groupRespnse,
	})

}
