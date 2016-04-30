package operations

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/nmonterroso/lolchest.win/models"
	"github.com/nmonterroso/lolchest.win/riotapi"
)

type getAllChampDataHandler struct {
	cache   []*models.ChampionData
	riotAPI riotapi.RiotApi
}

func NewGetAllChampDataHandler(api riotapi.RiotApi) GetAllChampDataHandler {
	return &getAllChampDataHandler{
		cache:   nil,
		riotAPI: api,
	}
}

func (h *getAllChampDataHandler) Handle() middleware.Responder {
	if h.cache == nil {
		h.cache = h.riotAPI.GetChampionData()
	}

	return NewGetAllChampDataOK().WithPayload(h.cache)
}
