package httputil

import (
	"net/http"
)

// 声明一个包级别的变量timeouts，类型是Timeout
var timeouts = DefaultTimeout

/**
 * @Description: 创建一个http services
 * @param handler:请求处理器，用于处理请求
 */
func NewHttpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:           handler,
		ReadTimeout:       DefaultTimeout.ReadTimeout,
		ReadHeaderTimeout: DefaultTimeout.ReadHeaderTimeout,
		WriteTimeout:      DefaultTimeout.WriteTimeout,
		IdleTimeout:       DefaultTimeout.IdleTimeout,
	}
}
