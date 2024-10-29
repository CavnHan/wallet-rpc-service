package httputil

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"sync/atomic"
)

type HTTPServer struct {
	listener net.Listener
	srv      *http.Server
	closed   atomic.Bool
}

// 定义一个函数类型HTTPOption，返回值是error
type HTTPOption func(sev *HTTPServer) error

/**
 * @Description: 创建一个http services
 * @param addr:监听地址
 * @param handler:请求处理器，用于处理请求
 * @param opts:可变响应参数
 */
func StartHttpServer(addr string, handler http.Handler, opts ...HTTPOption) (*HTTPServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("start http services failed, err:%v\n", err)
		return nil, errors.New("start http services failed")
	}
	srvCtx, srvCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			return srvCtx
		},
	}
	out := &HTTPServer{
		listener: listener,
		srv:      srv,
	}

	// 遍历opts，将opts中的函数应用到out上
	//可以进行一些自定义操作，比如设置http server的一些参数
	//进行动态的设置
	for _, opt := range opts {
		if err := opt(out); err != nil {
			srvCancel()
			fmt.Printf("services apply err:%v\n", err)
			return nil, errors.New("one of http op failed")
		}
	}
	go func() {
		// 启动http服务
		// Serve方法监听并接受传入的连接，然后为每个连接创建一个新的服务goroutine
		//这里会一直阻塞，直到http服务关闭
		err := out.srv.Serve(listener)            // 启动 HTTP 服务器并监听传入的连接
		srvCancel()                               // 取消服务器的上下文
		if errors.Is(err, http.ErrServerClosed) { // 检查错误是否是服务器正常关闭
			out.closed.Store(true) // 将服务器状态设置为已关闭
		} else {
			panic("unexpected error: " + err.Error()) // 抛出包含错误信息的 panic
		}
	}()
	return out, nil
}

// 检查关闭状态
func (hs *HTTPServer) Closed() bool {
	return hs.closed.Load()
}

// 关闭http services
func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	return hs.srv.Shutdown(ctx)
}

func (hs *HTTPServer) Close() error {
	return hs.srv.Close()
}

func (hs *HTTPServer) Addr() net.Addr {

	return hs.listener.Addr()
}
func WithMaxHeaderBytes(max int) HTTPOption {
	return func(server *HTTPServer) error {
		server.srv.MaxHeaderBytes = max
		return nil
	}
}

func (hs *HTTPServer) Stop(ctx context.Context) error {
	if err := hs.Shutdown(ctx); err != nil {
		if errors.Is(err, ctx.Err()) {
			return hs.Close()
		}
		return err
	}
	return nil
}
