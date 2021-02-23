package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"

	"zdmin/api/internal/config"
	"zdmin/api/internal/handler"
	"zdmin/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(
		func(err error) (int, interface{}) {
			if st, ok := status.FromError(err); ok {
				return http.StatusOK, map[string]interface{}{
					"code":    st.Code(),
					"message": st.Message(),
				}
			} else {
				return http.StatusInternalServerError, nil
			}
		},
	)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
