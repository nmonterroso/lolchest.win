package operations

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nmonterroso/lolchest.win/riotapi"
)

type getSummonerHandler struct {
	bridge riotapi.RiotApiBridge
}

func NewGetSummonerHandler(bridge riotapi.RiotApiBridge) GetSummonerHandler {
	return &getSummonerHandler{
		bridge: bridge,
	}
}

func (h *getSummonerHandler) Handle(params GetSummonerParams) middleware.Responder {
	return NewGetSummonerOK().WithPayload(h.bridge.GetSummonerData(params.Name))
}
