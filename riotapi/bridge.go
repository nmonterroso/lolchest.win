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
	params := operations.NewGetChampionDataParams().WithRegion(defaultRegion)
	data, err := client.Default.Operations.GetChampionData(params, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	var champions []*models.ChampionData
	urlBase := api.staticAssetURLBase()

	for _, champ := range data.Payload.Data {
		iconURL := fmt.Sprintf("%s/%s", urlBase, *champ.Image.Full)
		champions = append(champions, &models.ChampionData{
			IconURL: &iconURL,
			ID:      champ.ID,
			Name:    champ.Name,
		})
	}

	return champions
}

func (api *riotAPI) staticAssetURLBase() string {
	params := operations.NewGetStaticAssetVersionsParams().WithRegion(defaultRegion)
	data, err := client.Default.Operations.GetStaticAssetVersions(params, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return ""
	}

	return fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/champion", data.Payload[0])
}
