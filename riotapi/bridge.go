package riotapi

import (
	"fmt"
	"reflect"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/nmonterroso/lolchest.win/models"
	"github.com/nmonterroso/lolchest.win/riotapi/client"
	clientops "github.com/nmonterroso/lolchest.win/riotapi/client/operations"
)

type RiotApiBridge interface {
	GetChampionData() []*models.ChampionData
	GetSummonerData(name string) *models.Summoner
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

func (api *riotAPIBridge) GetSummonerData(name string) *models.Summoner {
	profileParams := clientops.NewGetSummonerProfileParams().WithSummonerNames(name)
	profileResponse, err := client.Default.Operations.GetSummonerProfile(profileParams, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	urlBase := api.staticAssetURLBase()
	var summoner *models.Summoner

	// TODO: there might be more than one, this will take just the last one
	for _, summonerProfile := range profileResponse.Payload {
		iconURL := fmt.Sprintf("%s/profileicon/%d.png", urlBase, *summonerProfile.ProfileIconID)
		summoner = &models.Summoner{
			ID:             summonerProfile.ID,
			Name:           summonerProfile.Name,
			ProfileIconURL: &iconURL,
			ChampData:      make([]*models.SummonerChampChestData, 0),
		}
	}

	masteryParams := clientops.NewGetSummonerChampionMasteryParams().WithSummonerID(*summoner.ID)
	masteryResponse, err := client.Default.Operations.GetSummonerChampionMastery(masteryParams, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil
	}

	for _, mastery := range masteryResponse.Payload {
		chestAvailable := !*mastery.ChestGranted
		summoner.ChampData = append(summoner.ChampData, &models.SummonerChampChestData{
			ChampID:          mastery.ChampionID,
			ChestIsAvailable: &chestAvailable,
			HighestGrade:     mastery.HighestGrade,
		})
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
