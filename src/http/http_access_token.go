package http

import (
	"bookstore-oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpires(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err.Error)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h *accessTokenHandler) UpdateExpires(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	var at access_token.AccessToken
	at.AccessToken = accessTokenId
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateExpires(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
