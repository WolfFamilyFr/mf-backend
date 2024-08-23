package comics

import (
	"context"
	"testing"

	"github.com/bchaillou003/marvel-family-backend/application"
	"github.com/bchaillou003/marvel-family-backend/application/mocks"
	models "github.com/bchaillou003/marvel-family-backend/models/comics"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_GetAllCharacter(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllCharacter", mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetAllCharacter(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetCharacterByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetCharacterByID", mock.Anything, mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetCharacterByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComicByCharacterID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllComicByCharacterID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComicByCharacterID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEventByCharacterID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllEventByCharacterID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEventByCharacterID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllSeriesByCharacterID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllSeriesByCharacterID", mock.Anything, mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetAllSeriesByCharacterID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStoryByCharacterID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllStoryByCharacterID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStoryByCharacterID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComic(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllComic", mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComic(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetComicByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetComicByID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetComicByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCharacterByComicID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCharacterByComicID", mock.Anything, mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetAllCharacterByComicID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCreatorByComicID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCreatorByComicID", mock.Anything, mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetAllCreatorByComicID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEventByComicID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllEventByComicID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEventByComicID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStoryByComicID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllStoryByComicID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStoryByComicID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCreator(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllCreator", mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetAllCreator(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetCreatorByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetCreatorByID", mock.Anything, mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetCreatorByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComicByCreatorID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllComicByCreatorID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComicByCreatorID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEventByCreatorID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllEventByCreatorID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEventByCreatorID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllSeriesByCreatorID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllSeriesByCreatorID", mock.Anything, mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetAllSeriesByCreatorID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStoryByCreatorID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllStoryByCreatorID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStoryByCreatorID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEvent(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllEvent", mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEvent(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetEventByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetEventByID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetEventByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCharacterByEventID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCharacterByEventID", mock.Anything, mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetAllCharacterByEventID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComicByEventID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllComicByEventID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComicByEventID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCreatorByEventID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCreatorByEventID", mock.Anything, mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetAllCreatorByEventID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllSeriesByEventID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllSeriesByEventID", mock.Anything, mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetAllSeriesByEventID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStoryByEventID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllStoryByEventID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStoryByEventID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllSeries(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllSeries", mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetAllSeries(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetSeriesByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetSeriesByID", mock.Anything, mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetSeriesByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCharacterBySeriesID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCharacterBySeriesID", mock.Anything, mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetAllCharacterBySeriesID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComicBySeriesID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllComicBySeriesID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComicBySeriesID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCreatorBySeriesID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCreatorBySeriesID", mock.Anything, mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetAllCreatorBySeriesID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEventBySeriesID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllEventBySeriesID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEventBySeriesID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStoryBySeriesID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllStoryBySeriesID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStoryBySeriesID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllStory(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetAllStory", mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetAllStory(ctx)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetStoryByID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	stories := models.StoryDataWrapper{
		Data: models.StoryDataContainer{
			Results: []models.Story{
				{
					ID: 1,
				},
			},
		},
	}

	sdk.On("GetStoryByID", mock.Anything, mock.Anything).Return(&stories, nil)

	// Exercise
	result, err := GetStoryByID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, stories.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCharacterByStoryID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	characters := models.CharacterDataWrapper{
		Data: models.CharacterDataContainer{
			Results: []models.Character{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCharacterByStoryID", mock.Anything, mock.Anything).Return(&characters, nil)

	// Exercise
	result, err := GetAllCharacterByStoryID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, characters.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllComicByStoryID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	comics := models.ComicDataWrapper{
		Data: models.ComicDataContainer{
			Results: []models.Comic{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllComicByStoryID", mock.Anything, mock.Anything).Return(&comics, nil)

	// Exercise
	result, err := GetAllComicByStoryID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, comics.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllCreatorByStoryID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	creators := models.CreatorDataWrapper{
		Data: models.CreatorDataContainer{
			Results: []models.Creator{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllCreatorByStoryID", mock.Anything, mock.Anything).Return(&creators, nil)

	// Exercise
	result, err := GetAllCreatorByStoryID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, creators.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllEventByStoryID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	events := models.EventDataWrapper{
		Data: models.EventDataContainer{
			Results: []models.Event{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllEventByStoryID", mock.Anything, mock.Anything).Return(&events, nil)

	// Exercise
	result, err := GetAllEventByStoryID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, events.Data.Results[0].ID, result.Data.Results[0].ID)
}

func Test_GetAllSeriesByStoryID(t *testing.T) {
	// Setup
	ctx := context.Background()
	app, _, sdk := mocks.NewMockedApplication()
	ctx = application.ToContext(ctx, app)

	series := models.SeriesDataWrapper{
		Data: models.SeriesDataContainer{
			Results: []models.Series{
				{
					ID: 2,
				},
			},
		},
	}

	sdk.On("GetAllSeriesByStoryID", mock.Anything, mock.Anything).Return(&series, nil)

	// Exercise
	result, err := GetAllSeriesByStoryID(ctx, 1)

	// Verify
	require.NoError(t, err)
	require.Equal(t, series.Data.Results[0].ID, result.Data.Results[0].ID)
}
