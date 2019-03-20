package gate

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

// RegisterRoute : registered route
func (s *Service) RegisterRoute(router, method string, handler module.Handler) {
	if _, ok := s.routeMap.Load(router); ok {
		logger.Sugar.Fatal("repeated router : %s", router)
	}

	s.routeMap.Store(router, module.Router{
		Method:  method,
		Handler: handler,
	})
}

func (s *Service) registerToGate(mux *http.ServeMux) {
	s.routeMap.Range(func(k, v interface{}) bool {
		router, err := parse.StringWithError(k)
		if err != nil {
			logger.Sugar.Fatalf("illegal http router[%v], not string, parse error [%v]", k, err)
		}
		handler := v.(module.Router)

		mux.HandleFunc(router, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == handler.Method {
				statusCode, resp := handler.Handler(r)
				data, err := proto.Marshal(resp)
				if err != nil {
					statusCode = http.StatusInternalServerError
				}
				w.WriteHeader(statusCode)
				w.Write(data)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
		})
		return true
	})
}