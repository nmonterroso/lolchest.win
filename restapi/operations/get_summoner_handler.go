package operations

import (
	"github.com/go-openapi/runtime"
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
	summonerData, err := h.bridge.GetSummonerData(params.Name)

	if err != nil {
		switch err.(type) {
		case *runtime.APIError:
			return NewGetSummonerInternalServerError()
		default:
			return NewGetSummonerBadGateway()
		}
	}

	return NewGetSummonerOK().WithPayload(summonerData)
}
