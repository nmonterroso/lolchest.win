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
	GetSummonerData(name string) (*models.Summoner, error)
}

type riotAPIBridge struct {
	auth runtime.ClientAuthInfoWriter
}

func NewRiotAPI(apiKey string) RiotApiBridge {
	return &riotAPIBridge{
		auth: httptransport.APIKeyAuth("api_key", "query", apiKey),
	}
}

func (api *riotAPIBridge) GetSummonerData(name string) (*models.Summoner, error) {
	urlBase, err := api.staticAssetURLBase()
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil, err
	}

	champions, err := api.getChampions(urlBase)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil, err
	}

	summoner, err := api.getSummoner(name, urlBase)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil, err
	}

	api.fillMasteries(*summoner.ID, champions)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return nil, err
	}

	for _, champ := range champions {
		summoner.ChampMastery = append(summoner.ChampMastery, champ)
	}

	return summoner, nil
}

func (api *riotAPIBridge) getChampions(iconURLBase string) (map[int64]*models.ChampionMastery, error) {
	resp, err := client.Default.Operations.GetChampionData(nil, api.auth)

	if err != nil {
		return nil, err
	}

	champions := make(map[int64]*models.ChampionMastery)
	for _, champ := range resp.Payload.Data {
		iconURL := fmt.Sprintf("%s/champion/%s", iconURLBase, *champ.Image.Full)
		chestAvailable := true

		champions[*champ.ID] = &models.ChampionMastery{
			ChampID:          champ.ID,
			ChampName:        champ.Name,
			ChampIconURL:     &iconURL,
			ChestIsAvailable: &chestAvailable,
		}
	}

	return champions, nil
}

func (api *riotAPIBridge) getSummoner(name string, urlBase string) (*models.Summoner, error) {
	params := clientops.NewGetSummonerProfileParams().WithSummonerNames(name)
	resp, err := client.Default.Operations.GetSummonerProfile(params, api.auth)

	if err != nil {
		return nil, err
	}

	var summoner *models.Summoner
	// TODO: there might be more than one, this will take just the last one
	for _, summonerProfile := range resp.Payload {
		iconURL := fmt.Sprintf("%s/profileicon/%d.png", urlBase, *summonerProfile.ProfileIconID)
		summoner = &models.Summoner{
			ID:             summonerProfile.ID,
			Name:           summonerProfile.Name,
			ProfileIconURL: &iconURL,
			ChampMastery:   make([]*models.ChampionMastery, 0),
		}
	}

	return summoner, nil
}

func (api *riotAPIBridge) fillMasteries(summonerID int64, champions map[int64]*models.ChampionMastery) error {
	params := clientops.NewGetSummonerChampionMasteryParams().WithSummonerID(summonerID)
	resp, err := client.Default.Operations.GetSummonerChampionMastery(params, api.auth)

	if err != nil {
		return err
	}

	for _, mastery := range resp.Payload {
		if champion, ok := champions[*mastery.ChampionID]; ok {
			chestAvailable := !*mastery.ChestGranted
			champion.ChestIsAvailable = &chestAvailable

			if mastery.HighestGrade != nil {
				champion.HighestGrade = *mastery.HighestGrade
			}
		}
	}

	return nil
}

func (api *riotAPIBridge) staticAssetURLBase() (string, error) {
	data, err := client.Default.Operations.GetStaticAssetVersions(nil, api.auth)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s %v", reflect.TypeOf(err), err))
		return "", err
	}

	return fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img", data.Payload[0]), nil
}
