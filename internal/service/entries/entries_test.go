package entries

import (
	"errors"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type mockEntriesRepo struct {
	entries []entry.Entry
	err     error
}

func (m *mockEntriesRepo) GetAll() ([]entry.Entry, error) {
	return m.entries, m.err
}

func (m *mockEntriesRepo) Create(e entry.Entry) error {
	if m.err != nil {
		return m.err
	}
	m.entries = append(m.entries, e)
	return nil
}

func TestEntries_GetAll_Success(t *testing.T) {
	mock := &mockEntriesRepo{
		entries: []entry.Entry{
			{Symbol: "BTCUSDT"},
		},
	}

	repo := &repository.Repository{
		Entries: mock,
	}

	svc := New(repo)

	res, err := svc.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res) != 1 || res[0].Symbol != "BTCUSDT" {
		t.Fatalf("unexpected result: %+v", res)
	}
}

func TestEntries_GetAll_Error(t *testing.T) {
	mock := &mockEntriesRepo{
		err: errors.New("some error"),
	}

	repo := &repository.Repository{
		Entries: mock,
	}

	svc := New(repo)

	_, err := svc.GetAll()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

