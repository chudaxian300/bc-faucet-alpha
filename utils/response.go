package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Date  interface{} `json:"date,omitempty"`
	Count int         `json:"count,omitempty"`
}

func (res *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, res)
	return
}
