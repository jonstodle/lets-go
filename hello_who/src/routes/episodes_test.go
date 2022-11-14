package routes

import (
	"encoding/json"
	"github.com/jonstodle/lets-go/hello_who/src/database/models"
	"github.com/jonstodle/lets-go/hello_who/src/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetEpisodes(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx, db, _ := test.NewContext(req, rec)

	var empty []models.Episode
	db.
		On("Select", &empty, mock.IsType(""), mock.IsType([]any{})).
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

	if a.NoError(getEpisodes(ctx)) {
		a.Equal(http.StatusOK, rec.Code)
		var episodes []models.Episode
		if a.NoError(json.NewDecoder(rec.Body).Decode(&episodes)) {
			a.Equal(1, len(episodes))
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

func TestGetEpisode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx, db, _ := test.NewContext(req, rec)
	ctx.SetPath("/episode/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	db.
		On(
			"Get",
			&models.Episode{},
			mock.MatchedBy(func(query string) bool { return strings.Contains(query, "id = ?") }),
			[]any{1}).
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

	if a.NoError(getEpisode(ctx)) {
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
