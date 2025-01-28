package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()
	var request posts.CreateCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIdStr := c.Param("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid post id",
		})
		return
	}
	userId := c.GetInt64("userId")
	err = h.postSvc.CreateComment(ctx, postId, userId, request)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
