package operations

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/nmonterroso/lolchest.win/models"
)

type getAllChampDataHandler struct {
	cache []*models.ChampionData
}

func NewGetAllChampDataHandler() GetAllChampDataHandler {
	return &getAllChampDataHandler{
		cache: nil,
	}
}

func (i *getAllChampDataHandler) Handle() middleware.Responder {
	if i.cache != nil {
		return NewGetAllChampDataOK().WithPayload(i.cache)
	}

}
