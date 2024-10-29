package httputil

import "net/http"

type WrappedResponseWriter struct {
	StatusCode  int
	ResponseLen int
	w           http.ResponseWriter
	wroteHeader bool
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{
		w:          w,
		StatusCode: 200,
	}
}

func (w *WrappedResponseWriter) Header() http.Header {
	return w.w.Header()
}
func (w *WrappedResponseWriter) WriteHeader(StatusCode int) {
	w.wroteHeader = true
	w.StatusCode = StatusCode
	w.w.WriteHeader(StatusCode)
}

func (w *WrappedResponseWriter) Write(bytes []byte) (int, error) {
	n, err := w.w.Write(bytes)
	w.ResponseLen += n
	return n, err
}
