package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.SignUp(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
