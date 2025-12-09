package handlers

import (
	"net/http"

	"L4.5/internal/model"
	"github.com/wb-go/wbf/ginext"
)

func (r *Router) StatsHandler(c *ginext.Context) {
	var nums model.Numbers

	err := c.BindJSON(&nums)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid JSON"})
		return
	}
	response := r.StatGetter.GetStats(nums)
	c.JSON(http.StatusOK, response)
}
