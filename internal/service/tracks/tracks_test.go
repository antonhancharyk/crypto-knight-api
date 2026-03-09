package tracks

import (
	"errors"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type mockTracksRepo struct {
	tracks []track.Track
	err    error
}

func (m *mockTracksRepo) GetAll(queryParams track.QueryParams) ([]track.Track, error) {
	return m.tracks, m.err
}

func (m *mockTracksRepo) Create(t track.Track) error {
	if m.err != nil {
		return m.err
	}
	m.tracks = append(m.tracks, t)
	return nil
}

func (m *mockTracksRepo) CreateBulk(tracks []track.Track) error {
	if m.err != nil {
		return m.err
	}
	m.tracks = append(m.tracks, tracks...)
	return nil
}

func (m *mockTracksRepo) GetAllHistory(queryParams track.QueryParams) ([]track.Track, error) {
	return m.tracks, m.err
}

func (m *mockTracksRepo) CreateBulkHistory(tracks []track.Track) error {
	if m.err != nil {
		return m.err
	}
	m.tracks = append(m.tracks, tracks...)
	return nil
}

func (m *mockTracksRepo) GetLastTracks() ([]track.Track, error) {
	return m.tracks, m.err
}

func TestTracks_GetAll_Success(t *testing.T) {
	mockTracks := &mockTracksRepo{
		tracks: []track.Track{
			{Symbol: "BTCUSDT", Interval: "1h"},
		},
	}

	repo := &repository.Repository{
		Tracks: mockTracks,
	}

	svc := New(repo)

	res, err := svc.GetAll(track.QueryParams{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res) != 1 || res[0].Symbol != "BTCUSDT" {
		t.Fatalf("unexpected result: %+v", res)
	}
}

func TestTracks_GetAll_Error(t *testing.T) {
	mockTracks := &mockTracksRepo{
		err: errors.New("some error"),
	}

	repo := &repository.Repository{
		Tracks: mockTracks,
	}

	svc := New(repo)

	_, err := svc.GetAll(track.QueryParams{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

