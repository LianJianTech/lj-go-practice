package middle

import (
	"bytes"
	"encoding/json"
	"github.com/LianJianTech/lj-go-common/errno"
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"io/ioutil"
	"lj-go-practice/handler"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()

		path := c.Request.URL.Path
		if path == "/api/check/health" || path == "/api/check/ram" || path == "/api/check/cpu" || path == "/api/check/disk" {
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, msg := -1, ""

		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "json.Unmarshal error, body:`%s`", blw.body.Bytes())
			code = errno.HandleError.Code
			msg = err.Error()
		} else {
			code = response.Code
			msg = response.Msg
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, msg)
	}
}
