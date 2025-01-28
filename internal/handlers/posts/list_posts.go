package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page index").Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page size").Error(),
		})
		return
	}
	ctx := c.Request.Context()
	response, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
