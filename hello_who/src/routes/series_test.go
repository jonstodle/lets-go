package routes

import (
	"encoding/json"
	"github.com/jonstodle/lets-go/hello_who/src/database/models"
	"github.com/jonstodle/lets-go/hello_who/src/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetSeriesEpisodes(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx, db, _ := test.NewContext(req, rec)
	ctx.SetPath("/series/:seriesnumber")
	ctx.SetParamNames("seriesnumber")
	ctx.SetParamValues("1")

	var episodes []models.Episode
	db.
		On(
			"Select",
			&episodes,
			mock.IsType(""),
			[]any{1}).
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]models.Episode)
			*arg = []models.Episode{
				{
					Id:                1,
					StoryNumber:       "175",
					SeriesNumber:      1,
					EpisodeNumber:     1,
					Title:             "The First",
					DoctorActor:       1,
					DirectedBy:        2,
					WrittenBy:         3,
					OriginalAirDate:   time.Date(2012, 1, 2, 3, 4, 5, 6, time.Local),
					Viewers:           1.4,
					AppreciationIndex: 87,
					CreatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
					UpdatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
				},
			}
		}).
		Once()

	a := assert.New(t)

	if a.NoError(getSeriesEpisodes(ctx)) {
		a.Equal(http.StatusOK, rec.Code)
		var episodes []models.Episode
		if a.NoError(json.NewDecoder(rec.Body).Decode(&episodes)) {
			a.Equal(models.Episode{
				Id:                1,
				StoryNumber:       "175",
				SeriesNumber:      1,
				EpisodeNumber:     1,
				Title:             "The First",
				DoctorActor:       1,
				DirectedBy:        2,
				WrittenBy:         3,
				OriginalAirDate:   time.Date(2012, 1, 2, 3, 4, 5, 6, time.Local),
				Viewers:           1.4,
				AppreciationIndex: 87,
				CreatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
				UpdatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
			}, episodes[0])
		}
	}
}

func TestGetSeriesEpisode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx, db, _ := test.NewContext(req, rec)
	ctx.SetPath("/series/:seriesNumber/episode/:episodeNumber")
	ctx.SetParamNames("seriesNumber", "episodeNumber")
	ctx.SetParamValues("1", "1")

	db.
		On(
			"Get",
			&models.Episode{},
			mock.IsType(""),
			[]any{1, 1}).
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Episode)
			*arg = models.Episode{
				Id:                1,
				StoryNumber:       "175",
				SeriesNumber:      1,
				EpisodeNumber:     1,
				Title:             "The First",
				DoctorActor:       1,
				DirectedBy:        2,
				WrittenBy:         3,
				OriginalAirDate:   time.Date(2012, 1, 2, 3, 4, 5, 6, time.Local),
				Viewers:           1.4,
				AppreciationIndex: 87,
				CreatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
				UpdatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
			}
		}).
		Once()

	a := assert.New(t)

	if a.NoError(getSeriesEpisode(ctx)) {
		a.Equal(http.StatusOK, rec.Code)
		var episode models.Episode
		if a.NoError(json.NewDecoder(rec.Body).Decode(&episode)) {
			a.Equal(models.Episode{
				Id:                1,
				StoryNumber:       "175",
				SeriesNumber:      1,
				EpisodeNumber:     1,
				Title:             "The First",
				DoctorActor:       1,
				DirectedBy:        2,
				WrittenBy:         3,
				OriginalAirDate:   time.Date(2012, 1, 2, 3, 4, 5, 6, time.Local),
				Viewers:           1.4,
				AppreciationIndex: 87,
				CreatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
				UpdatedAt:         time.Date(2022, 1, 2, 3, 4, 5, 6, time.Local),
			}, episode)
		}
	}
}
