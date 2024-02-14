package pokemon

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-htmx-demo/external/pokeapi"
)

type PokemonService interface {
	ResolvePokemonByFilter(ctx context.Context, filter PokemonFilter) (resp []PokemonResponse, err error)
	ResolvePokemonDetailByID(ctx context.Context, id int64) (resp PokemonDetail, err error)
}

type PokemonServiceImpl struct {
	PokeApiClient pokeapi.Client
}

func ProvidePokemonServiceImpl(pokeApiClient pokeapi.Client) *PokemonServiceImpl {
	return &PokemonServiceImpl{
		PokeApiClient: pokeApiClient,
	}
}

func (s *PokemonServiceImpl) ResolvePokemonByFilter(ctx context.Context, filter PokemonFilter) (resp []PokemonResponse, err error) {
	// TODO: set pagination

	if !filter.Page.Valid || filter.Page.Int64 == 0 {
		filter.SetDefaultPage()
	}

	if !filter.Limit.Valid || filter.Limit.Int64 == 0 {
		filter.SetDefaultLimit()
	}

	pokemons, err := s.PokeApiClient.GetPokemonList(pokeapi.PokemonListRequest{
		Limit:  filter.Limit.Int64,
		Offset: 0,
	})

	resp = NewPokemonResponseFromResult(pokemons.GetResult())

	return
}

func (s *PokemonServiceImpl) ResolvePokemonDetailByID(ctx context.Context, id int64) (resp PokemonDetail, err error) {
	pokemon, err := s.PokeApiClient.GetPokemonByID(id)
	if err != nil {
		log.Error().Err(err).Int("id", int(id)).Msg("[ResolvePokemonDetailByID]")
		return
	}

	resp = NewPokemonDetailResponseFromResult(pokemon)

	return
}
