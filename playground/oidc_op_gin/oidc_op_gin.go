package main

import (
	"context"
	"crypto/sha256"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/zitadel/oidc/example/server/storage"
	"github.com/zitadel/oidc/pkg/op"
	"golang.org/x/text/language"
	"ingress-authproxy/playground/oidc_op_gin/exampleop"
	"net/http"
	"os"
)

const pathLoggedOut = "/oidc/logged-out"

// newOP will create an OpenID Provider for localhost on a specified port with a given encryption key
// and a predefined default logout uri
// it will enable all options (see descriptions)
func newOP(ctx context.Context, storage op.Storage, issuer string, key [32]byte) (op.OpenIDProvider, error) {
	config := &op.Config{
		Issuer:    issuer,
		CryptoKey: key,

		// will be used if the end_session endpoint is called without a post_logout_redirect_uri
		DefaultLogoutRedirectURI: pathLoggedOut,

		// enables code_challenge_method S256 for PKCE (and therefore PKCE in general)
		CodeMethodS256: true,

		// enables additional client_id/client_secret authentication by form post (not only HTTP Basic Auth)
		AuthMethodPost: true,

		// enables additional authentication by using private_key_jwt
		AuthMethodPrivateKeyJWT: true,

		// enables refresh_token grant use
		GrantTypeRefreshToken: true,

		// enables use of the `request` Object parameter
		RequestObjectSupported: true,

		// this example has only static texts (in English), so we'll set the here accordingly
		SupportedUILocales: []language.Tag{language.English},
	}
	handler, err := op.NewOpenIDProvider(ctx, config, storage,
		// as an example on how to customize an endpoint this will change the authorization_endpoint from /authorize to /auth
		op.WithCustomAuthEndpoint(op.NewEndpoint("/auth")),
	)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

// setupLoginServer creates an OIDC server with Issuer=http://localhost:<port>
//
// Use one of the pre-made clients in storage/clients.go or register a new one.
func setupLoginServer(storage *storage.Storage, provider op.OpenIDProvider) *mux.Router {

	router := mux.NewRouter()

	// for simplicity, we provide a very small default page for users who have signed out
	router.HandleFunc(pathLoggedOut, func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("signed out successfully"))
		if err != nil {
			log.Printf("error serving logged out page: %v", err)
		}
	})

	// the provider will only take care of the OpenID Protocol, so there must be some sort of UI for the login process
	// for the simplicity of the example this means a simple page with username and password field
	l := exampleop.NewLogin(storage, op.AuthCallbackURL(provider))

	// regardless of how many pages / steps there are in the process, the UI must be registered in the router,
	// so we will direct all calls to /login to the login UI
	router.PathPrefix("/oidc/login/").Handler(http.StripPrefix("/oidc/login", l.Router))

	// we register the http handler of the OP on the root, so that the discovery endpoint (/.well-known/openid-configuration)
	// is served on the correct path
	//
	// if your issuer ends with a path (e.g. http://localhost:9998/custom/path/),
	// then you would have to set the path prefix (/custom/path/)
	router.PathPrefix("/").Handler(provider.HttpHandler())
	router.PathPrefix("/oidc/").Handler(provider.HttpHandler())

	return router
}

func setupOIDCServer(ctx context.Context, issuer string, storage op.Storage, encryptionKey string) op.OpenIDProvider {
	// this will allow us to use an issuer with http:// instead of https://
	_ = os.Setenv(op.OidcDevMode, "true")

	// the OpenID Provider requires a 32-byte key for (token) encryption
	// be sure to create a proper crypto random key and manage it securely!
	key := sha256.Sum256([]byte(encryptionKey))

	// for simplicity, we provide a very small default page for users who have signed out

	// creation of the OpenIDProvider with the just created in-memory Storage
	provider, err := newOP(ctx, storage, issuer, key)
	if err != nil {
		log.Fatal(err)
	}

	return provider
}

func main() {
	ctx := context.Background()
	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	repo := storage.NewStorage(storage.NewUserStore())
	port := "9998"

	provider := setupOIDCServer(ctx, "http://localhost:"+port+"/oidc", repo, "encryptionKey")
	router := setupLoginServer(repo, provider)

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
