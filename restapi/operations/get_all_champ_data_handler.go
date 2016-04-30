package operations

import (
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
	return NewGetAllChampDataOK().WithPayload(h.bridge.GetChampionData())
}
