package middleware

import (
	"api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := utils.Log()

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime) / time.Millisecond

		reqMethod := c.Request.Method

		reqUrl := c.Request.URL

		//header := c.Request.Header

		proto := c.Request.Proto

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		err := c.Err()

		body,_ := ioutil.ReadAll(c.Request.Body)

		logger.WithFields(logrus.Fields{
			"latency_time": latencyTime,
			"req_method": reqMethod,
			"req_url": reqUrl,
			//"header":header,
			"proto":proto,
			"status_code":statusCode,
			"client_ip":clientIP,
			"err":err,
			"body":body,
		}).Info()
	}
}
