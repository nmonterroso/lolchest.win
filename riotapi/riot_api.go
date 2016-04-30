package riotapi

import "github.com/nmonterroso/lolchest.win/models"

type RiotApi interface {
	GetChampionData() []*models.ChampionData
}

type riotAPI struct {
	apiKey string
}

func NewRiotApi(apiKey string) RiotApi {
	return &riotAPI{
		apiKey: apiKey,
	}
}

func (api *riotAPI) GetChampionData() []*models.ChampionData {

}
