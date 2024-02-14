package pokeapi

import (
	"github.com/guregu/null"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemons []Pokemon

type PokemonDetailResponse struct {
	ID      int64        `json:"id"`
	Name    string       `json:"name"`
	Stats   []StatDetail `json:"stats"`
	Height  int64        `json:"height"`
	Weight  int64        `json:"weight"`
	Sprites Sprites      `json:"sprites"`
}

type StatDetail struct {
	BaseStat int64 `json:"base_stat"`
	Stat     Stat  `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonListRequest struct {
	Limit  int64
	Offset int64
}

type PokemonListResponse struct {
	Count    int64       `json:"count"`
	Next     null.String `json:"next"`
	Previous null.String `json:"previous"`
	Results  Pokemons    `json:"results"`
}

func (p PokemonListResponse) GetResult() Pokemons {
	return p.Results
}

type Sprites struct {
	Other Other `json:"other"`
}

type Other struct {
	OfficialArtWork OfficialArtWork `json:"official-artwork"`
}

type OfficialArtWork struct {
	FrontDefault string `json:"front_default"`
}
