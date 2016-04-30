package riotapi

import (
	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/nmonterroso/lolchest.win/models"
	"github.com/nmonterroso/lolchest.win/riotapi/client"
	"github.com/nmonterroso/lolchest.win/riotapi/client/operations"
)

const (
	defaultRegion = "na"
)

type RiotApi interface {
	GetChampionData() []*models.ChampionData
}

type riotAPI struct {
	auth runtime.ClientAuthInfoWriter
}

func NewRiotApi(apiKey string) RiotApi {
	return &riotAPI{
		auth: httptransport.APIKeyAuth("apiKey", "query", apiKey),
	}
}

func (api *riotAPI) GetChampionData() []*models.ChampionData {
	params := operations.NewGetChampionDataParams().WithRegion(defaultRegion)
	data, err := client.Default.Operations.GetChampionData(params, api.auth)
}
