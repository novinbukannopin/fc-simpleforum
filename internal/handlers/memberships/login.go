package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipSvc.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.LoginResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, response)
}
