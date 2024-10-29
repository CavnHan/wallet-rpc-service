package httputil

import "time"

var DefaultTimeout = HTTPTimeouts{
	ReadTimeout:       10 * time.Second,
	ReadHeaderTimeout: 10 * time.Second,
	WriteTimeout:      10 * time.Second,
	IdleTimeout:       15 * time.Second,
}

type HTTPTimeouts struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func WithTimeouts(timeouts HTTPTimeouts) HTTPOption {
	return func(s *HTTPServer) error {
		s.srv.ReadTimeout = timeouts.ReadTimeout
		s.srv.ReadHeaderTimeout = timeouts.ReadHeaderTimeout
		s.srv.WriteTimeout = timeouts.WriteTimeout
		s.srv.IdleTimeout = timeouts.IdleTimeout
		return nil
	}
}
