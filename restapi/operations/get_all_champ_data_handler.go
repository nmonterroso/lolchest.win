package operations

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nmonterroso/lolchest.win/riotapi"
)

type getAllChampDataHandler struct {
	bridge riotapi.RiotApiBridge
}

func NewGetAllChampDataHandler(bridge riotapi.RiotApiBridge) GetAllChampDataHandler {
	return &getAllChampDataHandler{
		bridge: bridge,
	}
}

func (h *getAllChampDataHandler) Handle() middleware.Responder {
	champData, err := h.bridge.GetChampionData()

	if err != nil {
		switch err.(type) {
		case *runtime.APIError:
			return NewGetAllChampDataBadGateway()
		default:
			return NewGetAllChampDataInternalServerError()
		}
	}

	return NewGetAllChampDataOK().WithPayload(champData)
}
