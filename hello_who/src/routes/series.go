package routes

import (
	"database/sql"
	"errors"
	"github.com/jonstodle/lets-go/hello_who/src/database/models"
	"github.com/labstack/echo/v4"
)

func RegisterSeriesRoutes(group *echo.Group) {
	group.GET("/:seriesNumber", getSeriesEpisodes)
	group.GET("/:seriesNumber/episode/:episodeNumber", getSeriesEpisode)
}

type getSeriesEpisodesData struct {
	SeriesNumber int `param:"seriesNumber"`
}

func getSeriesEpisodes(c echo.Context) error {
	var data getSeriesEpisodesData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	var episodes []models.Episode
	err := getDB(c).Select(
		&episodes,
		`
		SELECT *
		FROM episodes
		WHERE series_number = ?
		`,
		data.SeriesNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Episode query failed", err)
	} else if episodes == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, episodes)
}

type getSeriesEpisodeData struct {
	SeriesNumber  int `param:"seriesNumber"`
	EpisodeNumber int `param:"episodeNumber"`
}

func getSeriesEpisode(c echo.Context) error {
	var data getSeriesEpisodeData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	var episode models.Episode
	err := getDB(c).Get(
		&episode,
		`
		SELECT *
		FROM episodes
		WHERE series_number = ? AND episode_number = ?
		LIMIT 1
		`,
		data.SeriesNumber,
		data.EpisodeNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Episode query failed", err)
	} else if episode.Id == 0 {
		return c.NoContent(404)
	}

	return c.JSON(200, episode)
}
