package web

import (
	"encoding/json"
	"github.com/i1kondratiuk/visitors-counter/application"
	"net/http"
)

func getUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := application.GetTinuser().GetUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	users, err := application.GetTinuser().GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	rawData, _ := c.GetRawData()
	data := struct {
		Name string `json:"name"`
	}{}
	err := json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = application.GetUsersCounter().AddUser(data.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func getUserMatch(c *gin.Context) {
}
