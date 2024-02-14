package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type Client struct {
	RestyClient *resty.Client
}

func ProviceClient() Client {
	restyClient := resty.New()
	return Client{
		RestyClient: restyClient,
	}
}

func (c *Client) GetPokemonByID(id int64) (pokemonResponse PokemonDetailResponse, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", id)

	resp, err := c.RestyClient.R().Get(url)
	if err != nil {
		log.Error().Err(err).Int("id", int(id)).Msg("[GetPokemonByID]")
		return
	}

	// TODO: handle status code
	
	if err = json.Unmarshal(resp.Body(), &pokemonResponse); err != nil {
		log.Error().Err(err).Interface("body", resp.Body()).Msg("[Unmarshal]")
		return
	}

	return
}

func (c *Client) GetPokemonList(request PokemonListRequest) (pokemonListResponse PokemonListResponse, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%d&offset=%d", request.Limit, request.Offset)

	resp, err := c.RestyClient.R().Get(url)
	if err != nil {
		log.Error().Err(err).Interface("request", request).Msg("[GetPokemonList]")
		return
	}

	// TODO: handle status code

	if err = json.Unmarshal(resp.Body(), &pokemonListResponse); err != nil {
		log.Error().Err(err).Interface("body", resp.Body()).Msg("[Unmarshal]")
		return
	}

	return
}
