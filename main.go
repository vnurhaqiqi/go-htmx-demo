package main

import (
	"github.com/vnurhaqiqi/go-htmx-demo/external/pokeapi"
	"github.com/vnurhaqiqi/go-htmx-demo/internal/handler"
	"github.com/vnurhaqiqi/go-htmx-demo/internal/pokemon"
	"github.com/vnurhaqiqi/go-htmx-demo/internal/server"
)

func main() {
	// external pokeapi client
	pokeApiClient := pokeapi.ProviceClient()
	// pokemon service
	pokemonService := pokemon.ProvidePokemonServiceImpl(pokeApiClient)
	// pokemon handler
	pokemonHandler := handler.ProvidePokemonHandler(pokemonService)
	// server
	serverProvider := server.ProvideServer(*pokemonHandler)
	serverProvider.Initialized()
}
