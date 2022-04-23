package handler

import (
	"net/http"

	"microservice/pkg/infrastructure/ports"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	service ports.ServicePort
}

func New(serv ports.ServicePort) *HttpHandler {
	return &HttpHandler{service: serv}
}

func (h HttpHandler) CreateStuff(c *gin.Context) {
	id, e := h.service.CreateStuff()
	if e != nil {
		c.JSON(e.Status, gin.H{
			"data": nil,
			"error": gin.H{
				"code":        e.Code,
				"description": e.Description,
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": id,
		},
		"error": nil,
	})
}
