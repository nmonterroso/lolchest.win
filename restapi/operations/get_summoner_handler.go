package operations

import (
	"fmt"
	"reflect"

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
		switch err.(type) {
		case *runtime.APIError:
			return NewGetSummonerInternalServerError()
		default:
			fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
			return NewGetSummonerBadGateway()
		}
	}

	return NewGetSummonerOK().WithPayload(summonerData)
}
