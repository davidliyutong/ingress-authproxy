package main

import (
	"context"
	"log"
	"net/http"

	"ingress-authproxy/playground/oidc_op_debug/exampleop"
	"ingress-authproxy/playground/oidc_op_debug/storage"
)

func main() {
	ctx := context.Background()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore())
	prefix := "/oidc"

	port := "9998"
	externalPort := "8080"
	router := exampleop.SetupServer(ctx, "http://localhost:"+externalPort+prefix, prefix, storage)

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
