// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 20 * time.Second
	_defaultWriteTimeout    = 20 * time.Second
	_defaultAddr            = "api.legchelife.ru:1007"
	_defaultShutdownTimeout = 3 * time.Second

	dev  = "dev"
	prod = "prod"
	test = "test"
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, build string, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start(build)

	return s
}

func (s *Server) start(build string) {
	switch build {
	case prod:
		go func() {
			s.notify <- s.server.ListenAndServeTLS(
				"/etc/ssl/certs/legchelife.ru.crt",
				"/etc/ssl/certs/legchelife.ru.key",
			)
			close(s.notify)
		}()
	case dev, test:
		go func() {
			s.notify <- s.server.ListenAndServe()
			close(s.notify)
		}()
	}
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
