package api

import (
	"net/http"
	"time"

	joingroup "finances.jordis.golang/application/members/users_groups/commands/join-group"
	getusergroups "finances.jordis.golang/application/members/users_groups/queries/getUserGroups"
	"github.com/gin-gonic/gin"
)

type GroupRespnsoDTO struct {
	Pk          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Secret      string    `json:"secret"`
	Created_by  string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (app *App) GetUserGroupsHandler(c *gin.Context) {
	userID, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}

	command := getusergroups.GetUserGroupsQuery{
		UserId: userID,
	}

	groups, err := getusergroups.GetUserGroupsQueryHandler(command, app.GroupsRepo, app.UsersGroupsRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var groupsResponse []GroupRespnsoDTO
	for _, group := range groups {
		groupRes := GroupRespnsoDTO{
			Pk:         group.Pk.Val,
			Name:       group.Name.Val,
			Secret:     group.Secret.Val,
			Created_by: group.Created_by.Val,
			CreatedAt:  group.CreatedAt,
			UpdatedAt:  group.UpdatedAt,
		}
		groupsResponse = append(groupsResponse, groupRes)
	}
	c.JSON(http.StatusOK, gin.H{
		"groups": groupsResponse,
	})

}

func (app *App) JoinGroup(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	var input struct {
		Secret string `json:"group_secret"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	command := joingroup.JoinGroupCommand{
		Secret: input.Secret,
		UserId: userId,
	}
	err = joingroup.JoinGroupCommandHandler(command, app.UsersGroupsRepo, app.GroupsRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "joined group succesfully",
	})

}
