package handler

import (
	"strconv"

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

// TODO: provide search pokemon
func (h *PokemonHandler) SearchPokemon(c *fiber.Ctx) error {
	var filter pokemon.PokemonFilter

	limit, _ := strconv.Atoi(c.Query("limit"))
	filter.Limit = null.IntFrom(int64(limit))

	return c.Render("pokemon_results", fiber.Map{"Results": "OK"})
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

func (h *PokemonHandler) Detail(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	resp, err := h.PokemonService.ResolvePokemonDetailByID(c.Context(), int64(id))
	if err != nil {
		log.Error().Err(err).Msg("[Detail]")
		return err
	}

	return c.Render("pokemon_detail", fiber.Map{"Result": resp})
}
