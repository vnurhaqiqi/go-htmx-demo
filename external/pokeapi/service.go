package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type Client struct {
	BaseURL     string
	RestyClient *resty.Client
}

func ProviceClient() Client {
	restyClient := resty.New()
	return Client{
		BaseURL:     "https://pokeapi.co",
		RestyClient: restyClient,
	}
}

func (c *Client) GetPokemonList(request PokemonListRequest) (pokemonListResponse PokemonListResponse, err error) {
	url := fmt.Sprintf("%s/api/v2/pokemon?limit=%d&offset=%d", c.BaseURL, request.Limit, request.Offset)

	resp, err := c.RestyClient.R().Get(url)
	if err != nil {
		log.Error().Err(err).Interface("request", request).Msg("[GetPokemonList]")
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error().Int("statusCode", resp.StatusCode()).Interface("body", resp.Body()).Msg("[GetPokemonList]")
		err = fmt.Errorf("invoked pokeapi")
		return
	}

	if err = json.Unmarshal(resp.Body(), &pokemonListResponse); err != nil {
		log.Error().Err(err).Interface("body", resp.Body()).Msg("[Unmarshal]")
		return
	}

	return
}

func (c *Client) GetPokemonByName(name string) (pokemonResponse PokemonDetailResponse, err error) {
	url := fmt.Sprintf("%s/api/v2/pokemon/%s", c.BaseURL, name)

	resp, err := c.RestyClient.R().Get(url)
	if err != nil {
		log.Error().Err(err).Str("name", name).Msg("[GetPokemonByName]")
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error().Int("statusCode", resp.StatusCode()).Interface("body", resp.Body()).Msg("[GetPokemonByName]")
		switch resp.StatusCode() {
		case http.StatusNotFound:
			err = fmt.Errorf("not found")
		default:
			err = fmt.Errorf("invoked pokeapi")
		}

		return
	}

	if err = json.Unmarshal(resp.Body(), &pokemonResponse); err != nil {
		log.Error().Err(err).Interface("body", resp.Body()).Msg("[Unmarshal]")
		return
	}

	return
}
