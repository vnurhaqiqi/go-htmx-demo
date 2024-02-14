package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-htmx-demo/internal/handler"
)

type Server struct {
	PokemonHandler handler.PokemonHandler
}

func ProvideServer(pokemonHandler handler.PokemonHandler) *Server {
	return &Server{
		PokemonHandler: pokemonHandler,
	}
}

func (s *Server) Initialized() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", s.PokemonHandler.Home)

	pokemon := app.Group("/pokemon")
	pokemon.Post("/search", s.PokemonHandler.SearchByName)
	pokemon.Get("/search/:name", s.PokemonHandler.GetByName)

	err := app.Listen(":3000")
	if err != nil {
		log.Error().Err(err).Msg("failed to start server...")
	}
}
