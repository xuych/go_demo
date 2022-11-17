package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenPagination(c *gin.Context) (keyword string, size int, page int) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	query := c.Query("keyword")

	if p == 0 {
		p = 1
	}
	if s == 0 {
		s = 10
	}
	// params := UserQueryParam{
	// 	page: p,
	// 	size: s,
	// 	keyword:    keyword,
	// }
	return query, p, s
}
