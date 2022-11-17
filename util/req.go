package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenPagination(c *gin.Context) (param UserQueryParam) {
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	q := c.Query("q")
	if p == 0 {
		p = 1
	}
	if s == 0 {
		s = 10
	}
	// params := UserQueryParam{
	// 	page: p,
	// 	size: s,
	// 	q:    q,
	// }
	params := UserQueryParam{s, p, q}
	return params
}
