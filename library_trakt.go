package minutes

import (
	"encoding/json"
	"strconv"

	trakt "github.com/42minutes/go-trakt"
)

// TraktLibrary is a read-only Trakt.tv library
type TraktLibrary struct {
	client *trakt.Client
}

// NewTraktLibrary returns a TraktLibrary
func NewTraktLibrary(client *trakt.Client) *TraktLibrary {
	return &TraktLibrary{
		client: client,
	}
}

// UpsertShow returns ErrNotImplemented
func (l *TraktLibrary) UpsertShow(show *Show) error {
	return ErrNotImplemented
}

// UpsertSeason returns ErrNotImplemented
func (l *TraktLibrary) UpsertSeason(season *Season) error {
	return ErrNotImplemented
}

// UpsertEpisode returns ErrNotImplemented
func (l *TraktLibrary) UpsertEpisode(episode *Episode) error {
	return ErrNotImplemented
}

// GetShow returns a Show
// or errors with ErrNotFound, or ErrInternalServer
func (l *TraktLibrary) GetShow(showID string) (*Show, error) {
	id, err := strconv.Atoi(showID)
	if err != nil {
		return nil, ErrInternalServer
	}
	show, result := l.client.Shows().One(id)
	if result.Err != nil {
		return nil, ErrInternalServer
	}
	// TODO: proper way to convert between structs
	sh := Show{}
	data, err := json.Marshal(show)
	if err != nil {
		return nil, ErrInternalServer
	}
	json.Unmarshal(data, &sh)

	return &sh, nil
}

// GetShows returns ErrNotImplemented
func (l *TraktLibrary) GetShows() ([]*Show, error) {
	return nil, ErrNotImplemented
}

// GetSeason returns a Season
// or errors with ErrNotFound, or ErrInternalServer
func (l *TraktLibrary) GetSeason(seasonID string) (*Season, error) {
	return nil, ErrNotImplemented
}

// GetSeasonsByShow returns all Seasons for a show
// or errors with ErrNotFound, or ErrInternalServer
func (l *TraktLibrary) GetSeasonsByShow(showID string) ([]*Season, error) {
	id, err := strconv.Atoi(showID)
	if err != nil {
		return nil, ErrInternalServer
	}

	seasons, result := l.client.Seasons().All(id)
	if result.Err != nil {
		return nil, ErrInternalServer
	}

	seasRs := []*Season{}
	for _, season := range seasons {
		se := Season{}
		data, err := json.Marshal(season)
		if err != nil {
			return nil, ErrInternalServer
		}
		json.Unmarshal(data, &se)
		seasRs = append(seasRs, &se)
	}
	return seasRs, nil
}

// GetSeasonByNumber returns a Season given a Show's ID and a Season number
// or errors with ErrNotFound, ErrMissingShow, or ErrInternalServer
// TODO not working and returs ErrNotImplemented
func (l *TraktLibrary) GetSeasonByNumber(showID string, seasonNumber int) (*Season, error) {
	id, err := strconv.Atoi(showID)
	if err != nil {
		return nil, ErrInternalServer
	}

	season, result := l.client.Seasons().ByNumber(id, seasonNumber)
	if result.Err != nil {
		return nil, ErrInternalServer
	}
	se := Season{}
	data, err := json.Marshal(season)
	if err != nil {
		return nil, ErrInternalServer
	}
	json.Unmarshal(data, &se)

	return nil, ErrNotImplemented
}

// GetEpisode returns ErrNotImplemented
func (l *TraktLibrary) GetEpisode(episodeID string) (*Episode, error) {
	return nil, ErrNotImplemented
}

// GetEpisodeByNumber returns an Episode  given a Show's ID a Season number
// and Episode's number
// or errors with ErrNotFound, ErrMissingShow, or ErrInternalServer
func (l *TraktLibrary) GetEpisodeByNumber(showID string, seasonNumber, episodeNumber int) (*Episode, error) {
	id, err := strconv.Atoi(showID)
	if err != nil {
		return nil, ErrInternalServer
	}

	episode, result := l.client.Episodes().OneBySeasonByNumber(
		id, seasonNumber, episodeNumber,
	)
	if result.Err != nil {
		return nil, ErrInternalServer
	}

	ep := Episode{}
	data, err := json.Marshal(episode)
	if err != nil {
		return nil, ErrInternalServer
	}
	json.Unmarshal(data, &ep)
	return &ep, nil
}

// GetEpisodesBySeason returns ErrNotImplemented
func (l *TraktLibrary) GetEpisodesBySeason(seasonID string) ([]*Episode, error) {
	return nil, ErrNotImplemented
}

// GetEpisodesBySeasonNumber returns all Shows for a show and season number
// or errors with ErrNotFound, or ErrInternalServer
func (l *TraktLibrary) GetEpisodesBySeasonNumber(showID string, seasonNumber int) ([]*Episode, error) {
	id, err := strconv.Atoi(showID)
	if err != nil {
		return nil, ErrInternalServer
	}

	episodes, result := l.client.Episodes().AllBySeason(id, seasonNumber)
	if result.Err != nil {
		return nil, ErrInternalServer
	}
	epsRs := []*Episode{}
	for _, episode := range episodes {
		ep := Episode{}
		data, err := json.Marshal(episode)
		if err != nil {
			return nil, ErrInternalServer
		}
		json.Unmarshal(data, &ep)
		epsRs = append(epsRs, &ep)
	}
	return epsRs, nil
}

// QueryShowsByTitle returns all Shows that match a partial title ordered
// by their probability
// or errors with ErrInternalServer
func (l *TraktLibrary) QueryShowsByTitle(title string) ([]*Show, error) {
	shows, result := l.client.Shows().Search(title)
	if result.Err != nil {
		return nil, ErrInternalServer
	}

	shRs := []*Show{}
	for _, show := range shows {
		sh := Show{}
		data, err := json.Marshal(show)
		if err != nil {
			return nil, ErrInternalServer
		}
		json.Unmarshal(data, &sh)
		shRs = append(shRs, &sh)
	}
	return shRs, nil
}