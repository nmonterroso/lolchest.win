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
	data, _ := client.Default.Operations.GetChampionData(params, api.auth)
	var champions []*models.ChampionData

	for _, champ := range data.Payload.Data {
		champions = append(champions, &models.ChampionData{
			IconURL: champ.Image.Full,
			ID:      champ.ID,
			Name:    champ.Name,
		})
	}
	return champions
}
