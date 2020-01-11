package controllers

import (
	"bookstore-user-api/domain"
	"bookstore-user-api/services"
	"bookstore-user-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("user id should be a number"))
		return
	}

	isPublic, err := strconv.ParseBool(c.GetHeader("X-Public"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	marshaledResult, marshalErr := result.Marshal(isPublic)
	if marshalErr != nil {
		c.JSON(marshalErr.Status, marshalErr)
		return
	}

	c.JSON(http.StatusOK, marshaledResult)

}

func UpdateUser(c *gin.Context) {
	var user domain.User

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("user id should be a number"))
		return
	}
	user.Id = userId
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.UpdateUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("user id should be a number"))
		return
	}

	getErr := services.DeleteUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, "successfully delete user!")

}

func Search(c *gin.Context) {
	status := c.Query("status")

	isPublic, err := strconv.ParseBool(c.GetHeader("X-Public"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	users, findErr := services.FindUserByStatus(status)
	if findErr != nil {
		c.JSON(findErr.Status, err)
		return
	}

	marshaledUsers, marshalErr := users.Marshal(isPublic)
	if marshalErr != nil {
		c.JSON(marshalErr.Status, marshalErr)
	}

	c.JSON(http.StatusOK, marshaledUsers)
}
