package riotapi

import (
	"errors"
	"fmt"
	"time"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/nmonterroso/lolchest.win/models"
	"github.com/nmonterroso/lolchest.win/riotapi/client"
	clientops "github.com/nmonterroso/lolchest.win/riotapi/client/operations"
	clientmodels "github.com/nmonterroso/lolchest.win/riotapi/models"
)

type RiotApiBridge interface {
	GetSummonerData(region string, name string, refresh bool) (*models.Summoner, error)
}

type riotAPIBridge struct {
	auth    runtime.ClientAuthInfoWriter
	clients map[string]*client.Riot
	cache   Cache
}

var regions = []string{
	"br",
	"eune",
	"euw",
	"jp",
	"kr",
	"lan",
	"las",
	"na",
	"oce",
	"ru",
	"tr",
}

func NewRiotAPI(apiKey string) RiotApiBridge {
	clients := make(map[string]*client.Riot)
	for _, r := range regions {
		clients[r] = client.New(
			httptransport.New(fmt.Sprintf("%s.api.pvp.net", r), "/", []string{"https"}),
			strfmt.Default)
	}

	return &riotAPIBridge{
		auth:    httptransport.APIKeyAuth("api_key", "query", apiKey),
		clients: clients,
		cache:   NewCache(),
	}
}

func (api *riotAPIBridge) GetSummonerData(region string, name string, refresh bool) (*models.Summoner, error) {
	urlBase, err := api.staticAssetURLBase(region)
	if err != nil {
		return nil, err
	}

	champions, err := api.getChampions(region, urlBase)
	if err != nil {
		return nil, err
	}

	summoner, err := api.getSummoner(region, name, urlBase)
	if err != nil {
		return nil, err
	}

	api.fillMasteries(region, *summoner.ID, champions)
	if err != nil {
		return nil, err
	}

	for _, champ := range champions {
		summoner.ChampMastery = append(summoner.ChampMastery, champ)
	}

	return summoner, nil
}

func (api *riotAPIBridge) getChampions(region string, iconURLBase string) (map[int64]*models.ChampionMastery, error) {
	data, err := api.cache.GetOrSet(fmt.Sprintf("championList-%s", region), 24*time.Hour, func() (interface{}, error) {
		params := clientops.NewGetChampionDataParams().WithRegion(region)
		resp, err := client.Default.Operations.GetChampionData(params, api.auth)

		if err != nil {
			return nil, err
		}

		return resp.Payload.Data, nil
	})

	if err != nil {
		return nil, err
	}

	championList := data.(map[string]clientmodels.ChampionDto)
	champions := make(map[int64]*models.ChampionMastery)

	for _, champ := range championList {
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

func (api *riotAPIBridge) getSummoner(region string, name string, urlBase string) (*models.Summoner, error) {
	data, err := api.cache.GetOrSet(fmt.Sprintf("summoner-%s-%s", region, name), 5*24*time.Hour, func() (interface{}, error) {
		params := clientops.NewGetSummonerProfileParams().WithSummonerNames(name).WithRegion(region)
		resp, err := api.clientFor(region).Operations.GetSummonerProfile(params, api.auth)

		// TODO: cache something on 404?
		if err != nil {
			return nil, err
		}

		// TODO: there might be more than one, this will take just the last one
		for _, summonerProfile := range resp.Payload {
			return summonerProfile, nil
		}

		return nil, errors.New("no profiles returned")
	})

	if err != nil {
		return nil, err
	}

	summonerProfile := data.(clientmodels.SummonerDto)
	iconURL := fmt.Sprintf("%s/profileicon/%d.png", urlBase, *summonerProfile.ProfileIconID)

	return &models.Summoner{
		ID:             summonerProfile.ID,
		Name:           summonerProfile.Name,
		ProfileIconURL: &iconURL,
		ChampMastery:   make([]*models.ChampionMastery, 0),
	}, nil
}

func (api *riotAPIBridge) fillMasteries(region string, summonerID int64, champions map[int64]*models.ChampionMastery) error {
	platformID := regionToPlatformID(region)
	data, err := api.cache.GetOrSet(fmt.Sprintf("masteries-%s-%d", platformID, summonerID), 3*time.Hour, func() (interface{}, error) {
		params := clientops.NewGetSummonerChampionMasteryParams().WithSummonerID(summonerID).WithPlatformID(platformID)
		resp, err := api.clientFor(region).Operations.GetSummonerChampionMastery(params, api.auth)

		if err != nil {
			return nil, err
		}

		return resp.Payload, nil
	})

	if err != nil {
		return err
	}

	masteries := data.([]*clientmodels.ChampionMasteryDto)
	for _, mastery := range masteries {
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

func (api *riotAPIBridge) staticAssetURLBase(region string) (string, error) {
	data, err := api.cache.GetOrSet(fmt.Sprintf("staticAssetURL"), 24*time.Hour, func() (interface{}, error) {
		params := clientops.NewGetStaticAssetVersionsParams().WithRegion(region)
		resp, err := client.Default.Operations.GetStaticAssetVersions(params, api.auth)

		if err != nil {
			return "", err
		}

		return resp.Payload[0], nil
	})

	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img", data.(string)), nil
}

func (api *riotAPIBridge) clientFor(region string) *client.Riot {
	return api.clients[region] // validated by swagger spec, kinda ghetto but meh
}

func regionToPlatformID(region string) string {
	switch region {
	case "oce":
		return "oc1"
	case "eune":
		return "eun1"
	case "kr", "ru":
		return region
	case "lan":
		return "lan1"
	case "las":
		return "lan2"
	}

	return fmt.Sprintf("%s1", region) //TODO: there is LAN1 and LAN2 O.o
}
