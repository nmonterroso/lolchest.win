package restapi

import (
	"os"
	"crypto/tls"
	"net/http"
	"github.com/nmonterroso/lolchest.win/riotapi"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/nmonterroso/lolchest.win/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.LolchestWinAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LolchestWinAPI) http.Handler {
	apiKey := os.Getenv("RIOT_API_KEY")
	riotAPI := riotapi.NewRiotApi(apiKey)
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetAllChampDataHandler = operations.NewGetAllChampDataHandler(riotAPI)
	api.GetSummonerHandler = operations.NewGetSummonerHandler(riotAPI)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
