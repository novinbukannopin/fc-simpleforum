package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"net/http"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, c.GetInt64("userId"), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, memberships.RefreshResponse{RefreshToken: accessToken})

}
