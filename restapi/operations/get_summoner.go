package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSummonerHandlerFunc turns a function with the right signature into a get summoner handler
type GetSummonerHandlerFunc func(GetSummonerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSummonerHandlerFunc) Handle(params GetSummonerParams) middleware.Responder {
	return fn(params)
}

// GetSummonerHandler interface for that can handle valid get summoner params
type GetSummonerHandler interface {
	Handle(GetSummonerParams) middleware.Responder
}

// NewGetSummoner creates a new http.Handler for the get summoner operation
func NewGetSummoner(ctx *middleware.Context, handler GetSummonerHandler) *GetSummoner {
	return &GetSummoner{Context: ctx, Handler: handler}
}

/*GetSummoner swagger:route GET /{region}/{name} getSummoner

get data for a summoner

*/
type GetSummoner struct {
	Context *middleware.Context
	Handler GetSummonerHandler
}

func (o *GetSummoner) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetSummonerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
