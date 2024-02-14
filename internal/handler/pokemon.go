package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null"
	"github.com/rs/zerolog/log"

	// "github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-htmx-demo/internal/pokemon"
)

type PokemonHandler struct {
	PokemonService pokemon.PokemonService
}

func ProvidePokemonHandler(pokemonService pokemon.PokemonService) *PokemonHandler {
	return &PokemonHandler{
		PokemonService: pokemonService,
	}
}

func (h *PokemonHandler) Home(c *fiber.Ctx) error {
	resp, err := h.PokemonService.ResolvePokemonByFilter(c.Context(), pokemon.PokemonFilter{
		Limit: null.IntFrom(100),
	})

	if err != nil {
		log.Error().Err(err).Msg("[Home]")
		return err
	}

	return c.Render("index", fiber.Map{"Results": resp})
}

func (h *PokemonHandler) SearchByName(c *fiber.Ctx) error {
	name := c.FormValue("name")

	resp, err := h.PokemonService.ResolvePokemonDetailByName(c.Context(), name)
	if err != nil {
		log.Error().Err(err).Msg("[SearchByName]")
		return c.Render("pokemon_not_found", fiber.Map{})
	}

	return c.Render("pokemon_detail", fiber.Map{"Result": resp})
}

func (h *PokemonHandler) GetByName(c *fiber.Ctx) error {
	name := c.Params("name")

	resp, err := h.PokemonService.ResolvePokemonDetailByName(c.Context(), name)
	if err != nil {
		log.Error().Err(err).Msg("[SearchByName]")
		return err
	}

	return c.Render("pokemon_detail", fiber.Map{"Result": resp})
}
