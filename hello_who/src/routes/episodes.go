package routes

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jonstodle/lets-go/hello_who/src/database"
	"github.com/jonstodle/lets-go/hello_who/src/database/models"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func RegisterEpisodesRoutes(group *echo.Group) {
	group.GET("", getEpisodes)
	group.GET("/:id", getEpisode)
	group.GET("/onthisday", getEpisodesOnThisDay)
	group.GET("/onthisday/:id", getEpisodesOnThisDayResult)
}

func getEpisodes(c echo.Context) error {
	var episodes []models.Episode
	err := getDB(c).Select(
		&episodes,
		`
		SELECT *
		FROM episodes
		`)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Episode query failed", err)
	} else if episodes == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, episodes)
}

type getEpisodeData struct {
	Id int `param:"id"`
}

func getEpisode(c echo.Context) error {
	var data getEpisodeData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	var episode models.Episode
	err := getDB(c).Get(
		&episode,
		`
		SELECT *
		FROM episodes
		WHERE id = ?
		LIMIT 1
		`,
		data.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Episode query failed", err)
	} else if episode.Id == 0 {
		return c.NoContent(404)
	}

	return c.JSON(200, episode)
}

// =============================================
//
//	Async response
//
// =============================================
func getEpisodesOnThisDay(c echo.Context) error {
	today, err := getClock(c).Today("Europe/Oslo")
	if err != nil {
		return echo.NewHTTPError(500, "Failed to get today's date", err)
	}

	id := uuid.New()
	go runGetEpisodesOnDay(id, today, getDB(c))

	return c.JSON(200, map[string]any{
		"id": id,
	})
}

var store = make(map[uuid.UUID]struct {
	Episodes []models.Episode
	Error    error
})

func runGetEpisodesOnDay(id uuid.UUID, day time.Time, db database.Database) {
	log.Println("Starting heavy work...")
	defer log.Println("Finished heavy work")

	// Simulate a heavy work
	time.Sleep(10 * time.Second)

	var episodes []models.Episode
	err := db.Select(
		&episodes,
		`
		SELECT *
		FROM episodes
		WHERE STRFTIME('%m-%d', original_air_date) = STRFTIME('%m-%d', ?)
		`,
		day)

	store[id] = struct {
		Episodes []models.Episode
		Error    error
	}{episodes, err}
}

type getEpisodesOnThisDayResultData struct {
	Id uuid.UUID `param:"id"`
}

func getEpisodesOnThisDayResult(c echo.Context) error {
	var data getEpisodesOnThisDayResultData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	result := store[data.Id]
	if errors.Is(result.Error, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if result.Error != nil {
		return echo.NewHTTPError(500, "Episode query failed", result.Error)
	} else if result.Episodes == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, result.Episodes)
}
