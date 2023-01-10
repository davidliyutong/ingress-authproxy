package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/zitadel/oidc/example/server/storage"
	"ingress-authproxy/playground/oidc_op_gin/exampleop"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore())

	port := "9998"
	router := exampleop.SetupServer(ctx, r, "http://localhost:"+port, "/oidc", storage)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
