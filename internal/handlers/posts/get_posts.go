package posts

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {
	ctx := c.Request.Context()
	postIdStr := c.Param("postId")
	fmt.Println("postIdStr", postIdStr)
	fmt.Println("id", c.GetInt64("userId"))
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": errors.New("invalid post id").Error(),
		})
		return
	}

	response, err := h.postSvc.GetPostById(ctx, postId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
