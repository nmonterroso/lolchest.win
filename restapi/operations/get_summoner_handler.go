package operations

import (
	"fmt"

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
	refresh := false
	if params.Refresh != nil {
		refresh = *params.Refresh
	}

	summonerData, err := h.bridge.GetSummonerData(params.Region, params.Name, refresh)

	if err != nil {
		switch e := err.(type) {
		case *runtime.APIError:
			fmt.Println(fmt.Sprintf("%s(%s, %s) - %d", e.OperationName, params.Region, params.Name, e.Code))
			return NewGetSummonerInternalServerError()
		default:
			fmt.Println(fmt.Sprintf("%s (%s, %s)", e.Error(), params.Region, params.Name))
			return NewGetSummonerBadGateway()
		}
	}

	return NewGetSummonerOK().WithPayload(summonerData)
}
