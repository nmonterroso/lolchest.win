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
	refresh := false
	if params.Refresh != nil {
		refresh = *params.Refresh
	}

	summonerData, err := h.bridge.GetSummonerData(params.Region, params.Name, refresh)

	if err != nil {
		switch *err.Code {
		case 500:
			return NewGetSummonerInternalServerError()
		default:
			return NewGetSummonerBadGateway().WithPayload(err)
		}
	}

	return NewGetSummonerOK().WithPayload(summonerData)
}
