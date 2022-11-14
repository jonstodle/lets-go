package routes

import (
	"database/sql"
	"errors"
	"github.com/jonstodle/lets-go/hello_who/src/database/models"
	"github.com/labstack/echo/v4"
)

func RegisterPersonsRoutes(group *echo.Group) {
	group.GET("", getPersons)
	group.GET("/:id", getPerson)
	group.GET("/:id/episodes", getPersonEpisodes)
}

func getPersons(c echo.Context) error {
	var persons []models.Person
	err := getDB(c).Select(
		&persons,
		`
		SELECT *
		FROM persons
		`)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Persons query failed", err)
	} else if persons == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, persons)
}

type getPersonData struct {
	Id int `param:"id"`
}

func getPerson(c echo.Context) error {
	var data getPersonData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	var person models.Person
	err := getDB(c).Get(
		&person,
		`
		SELECT *
		FROM persons
		WHERE id = ?
		LIMIT 1
		`,
		data.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Person query failed", err)
	} else if person.Id == 0 {
		return c.NoContent(404)
	}

	return c.JSON(200, person)
}

type getPersonEpisodesData struct {
	Id int `param:"id"`
}

func getPersonEpisodes(c echo.Context) error {
	var data getPersonEpisodesData
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, "Bad input data", err)
	}

	var episodes []models.Episode
	err := getDB(c).Select(
		&episodes,
		`
		SELECT *
		FROM episodes
		WHERE ? in (doctor_actor, written_by, directed_by)
		`,
		data.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.NoContent(404)
	} else if err != nil {
		return echo.NewHTTPError(500, "Episode query failed", err)
	} else if episodes == nil {
		return c.NoContent(404)
	}

	return c.JSON(200, episodes)
}
