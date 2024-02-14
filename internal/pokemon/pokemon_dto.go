package pokemon

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/guregu/null"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-htmx-demo/external/pokeapi"
)

type PokemonResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *PokemonResponse) SetIDFromUrl() {
	parsedURL, err := url.Parse(p.Url)
	if err != nil {
		log.Warn().Err(err).Str("url", p.Url).Msg("error parse url")
	}
	// extract path from url
	path := parsedURL.Path

	// split path into segments
	segments := strings.Split(path, "/")

	id, _ := strconv.Atoi(segments[len(segments)-2])
	p.ID = int64(id)
}

type PokemonDetail struct {
	ID     int64        `json:"id"`
	Name   string       `json:"name"`
	Stats  []StatDetail `json:"stats"`
	Height int64        `json:"height"`
	Weight int64        `json:"weight"`
}

type StatDetail struct {
	BaseStat int64 `json:"base_stat"`
	Stat     Stat  `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonFilter struct {
	Page  null.Int
	Limit null.Int
}

func (f *PokemonFilter) SetDefaultPage() {
	f.Page = null.IntFrom(1)
}

func (f *PokemonFilter) SetDefaultLimit() {
	f.Limit = null.IntFrom(10)
}

func NewPokemonResponseFromResult(result pokeapi.Pokemons) []PokemonResponse {
	var pokemonResponses []PokemonResponse

	for _, res := range result {
		pokemon := PokemonResponse{
			Name: res.Name,
			Url:  res.Url,
		}
		pokemon.SetIDFromUrl()

		pokemonResponses = append(pokemonResponses, pokemon)
	}

	return pokemonResponses
}

func NewPokemonDetailResponseFromResult(result pokeapi.PokemonDetailResponse) PokemonDetail {
	var stats []StatDetail

	for _, resStat := range result.Stats {
		stat := Stat{Name: resStat.Stat.Name}
		stats = append(stats, StatDetail{
			BaseStat: resStat.BaseStat,
			Stat:     stat,
		})
	}

	return PokemonDetail{
		ID:     result.ID,
		Name:   result.Name,
		Height: result.Height,
		Weight: result.Weight,
		Stats:  stats,
	}
}
