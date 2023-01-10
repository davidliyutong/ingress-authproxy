package oidc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"ingress-authproxy/internal/apiserver/oidc/model"
	"ingress-authproxy/internal/config"
	"net/http"
	"strconv"

	"ingress-authproxy/internal/apiserver/oidc/exampleop"
	"ingress-authproxy/internal/apiserver/oidc/storage"
)

func RunOIDCServer(opt *config.OIDCOpt) {
	ctx := context.Background()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	s := storage.NewStorage(model.NewUserStore())
	prefix := "/oidc"

	router := exampleop.SetupServer(ctx, opt.BaseURL+opt.Prefix, prefix, s)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(int(opt.Port)),
		Handler: router,
	}
	log.Infof("[OIDC Server] Serving OIDC at port %v", opt.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
