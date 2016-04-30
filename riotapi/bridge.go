package riotapi

import (
	"fmt"
	"reflect"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/nmonterroso/lolchest.win/models"
	serverops "github.com/nmonterroso/lolchest.win/restapi/operations"
	"github.com/nmonterroso/lolchest.win/riotapi/client"
	clientops "github.com/nmonterroso/lolchest.win/riotapi/client/operations"
)

type RiotApiBridge interface {
	GetChampionData() []*models.ChampionData
	GetSummonerData(params serverops.GetSummonerParams) *models.Summoner
}

type riotAPIBridge struct {
	auth runtime.ClientAuthInfoWriter
}

func NewRiotApi(apiKey string) RiotApiBridge {
	return &riotAPIBridge{
		auth: httptransport.APIKeyAuth("api_key", "query", apiKey),
	}
}

func (api *riotAPIBridge) GetChampionData() []*models.ChampionData {
	data, err := client.Default.Operations.GetChampionData(nil, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	var champions []*models.ChampionData
	urlBase := api.staticAssetURLBase()

	for _, champ := range data.Payload.Data {
		iconURL := fmt.Sprintf("%s/champion/%s", urlBase, *champ.Image.Full)
		champions = append(champions, &models.ChampionData{
			IconURL: &iconURL,
			ID:      champ.ID,
			Name:    champ.Name,
		})
	}

	return champions
}

func (api *riotAPIBridge) GetSummonerData(params serverops.GetSummonerParams) *models.Summoner {
	params := clientops.NewGetSummonerProfileParams().WithSummonerNames(params.Name)
	data, err := client.Default.Operations.GetSummonerProfile(nil, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	urlBase := api.staticAssetURLBase()
	var summoner *models.Summoner

	// TODO: there might be more than one, this will take just the last one
	for _, summonerProfile := range data.Payload {
		summoner = &models.Summoner{
			ID: summonerProfile.ID,
			Name: summonerProfile.Name
			ProfileIconURL: fmt.Sprintf("%s/profileicon/%d.png", urlBase, summonerProfile.ProfileIconID)
			ChampData: make([]*models.SummonerChampChestData, 0)
		}
	}

	params = clientops.NewGetSummonerChampionMasteryParams().WithSummonerID(*summoner.ID)
	data, err = client.Default.Operations.GetSummonerChampionMastery(params, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	for mastery := range data.Payload {
		summoner.ChampData = append(summoner.ChampData, &models.SummonerChampChestData{
			ChampID: mastery.ChampionID,
			ChestIsAvailable: !mastery.ChestGranted,
			HighestGrade: mastery.HighestGrade,
		}
	}

	return summoner
}

func (api *riotAPIBridge) staticAssetURLBase() string {
	data, err := client.Default.Operations.GetStaticAssetVersions(nil, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return ""
	}

	return fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img", data.Payload[0])
}
