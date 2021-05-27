package requestid

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	prefix = ""
)

const (
	loggerFieldName    = "logger"
	requestIDFieldName = "requestID"
)

func init() {
	prefix, _ = os.Hostname()
	if prefix == "" {
		prefix = "localhost"
	}
}

func GetLogger(c *gin.Context) *logrus.Entry {
	GetRequestID(c) // 保证requestid存在
	value, _ := c.Get(loggerFieldName)
	return value.(*logrus.Entry)
}

func GetRequestID(c *gin.Context) string {
	requestID := c.GetString(requestIDFieldName)
	if requestID == "" {
		requestID = fmt.Sprintf("%s-%s", prefix, uuid.New().String())
		c.Set(requestIDFieldName, requestID)

		// FIXME
		logger := logrus.WithField(requestIDFieldName, requestID)
		c.Set(loggerFieldName, logger)
		c.Writer.Header().Set("Request-ID", requestID)
		return requestID
	}
	return requestID
}
