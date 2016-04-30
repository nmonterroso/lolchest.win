package riotapi

import (
	"fmt"
	"reflect"

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
		auth: httptransport.APIKeyAuth("api_key", "query", apiKey),
	}
}

func (api *riotAPI) GetChampionData() []*models.ChampionData {
	champData := "image"
	params := operations.NewGetChampionDataParams().WithRegion(defaultRegion).WithChampData(&champData)
	data, err := client.Default.Operations.GetChampionData(params, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

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
