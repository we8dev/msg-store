package httpserver

import (
	"net/http"
	"time"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
	_defaultAddr         = ":80"
	//_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	server *http.Server
	//notify          chan error
	//shutdownTimeout time.Duration
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server: httpServer,
		//notify:          make(chan error, 1),
		//shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	//s.start()

	return s
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

// TODO: Восстановить работу функций
//
//func (s *Server) start() {
//	go func() {
//		s.notify <- s.httpserver.ListenAndServe()
//		close(s.notify)
//	}()
//}

//func (s *Server) Notify() <-chan error {
//	return s.notify
//}

// TODO: Добавить метод для отключнеия сервера
//func (s *Server) Shutdown() error {
//	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
//	defer cancel()
//
//	return s.httpserver.Shutdown(ctx)
//}
