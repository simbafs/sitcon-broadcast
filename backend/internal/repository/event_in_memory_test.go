package repository

// TODO: 這個和 event_ent_test.go 測試感覺可以整合在一起，之後在說

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupEventRepo() (*EventInMemory, context.Context) {
	return NewEventInMemory(), context.Background()
}

func TestEventInMemory_Create_And_Get(t *testing.T) {
	repo, ctx := setupEventRepo()

	ev, err := repo.Create(ctx, "test", "https://example.com", "console.log('hi')")
	require.NoError(t, err)
	require.Equal(t, "test", ev.Name())
	require.Equal(t, "https://example.com", ev.URL())
	require.Equal(t, "console.log('hi')", ev.Script())

	got, err := repo.Get(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, ev, got)
}

func TestEventInMemory_Create_Duplicate(t *testing.T) {
	repo, ctx := setupEventRepo()
	_, _ = repo.Create(ctx, "test", "url", "script")

	_, err := repo.Create(ctx, "test", "url2", "script2")
	require.ErrorIs(t, err, ErrEventAlreadyExists)
}

func TestEventInMemory_Get_NotFound(t *testing.T) {
	repo, ctx := setupEventRepo()
	_, err := repo.Get(ctx, "notfound")
	require.ErrorIs(t, err, ErrCannotGetEvent)
}

func TestEventInMemory_List(t *testing.T) {
	repo, ctx := setupEventRepo()
	repo.Create(ctx, "a", "urlA", "scriptA")
	repo.Create(ctx, "b", "urlB", "scriptB")

	list, err := repo.List(ctx)
	require.NoError(t, err)
	require.Len(t, list, 2)
}

func TestEventInMemory_Update(t *testing.T) {
	repo, ctx := setupEventRepo()
	repo.Create(ctx, "target", "old-url", "old-script")

	err := repo.Update(ctx, "target", "new-url", "new-script")
	require.NoError(t, err)

	ev, _ := repo.Get(ctx, "target")
	require.Equal(t, "new-url", ev.URL())
	require.Equal(t, "new-script", ev.Script())
}

func TestEventInMemory_Update_NotFound(t *testing.T) {
	repo, ctx := setupEventRepo()

	err := repo.Update(ctx, "ghost", "url", "script")
	require.ErrorIs(t, err, ErrCannotGetEvent)
}

func TestEventInMemory_Delete(t *testing.T) {
	repo, ctx := setupEventRepo()
	repo.Create(ctx, "del", "url", "script")

	err := repo.Delete(ctx, "del")
	require.NoError(t, err)

	_, err = repo.Get(ctx, "del")
	require.ErrorIs(t, err, ErrCannotGetEvent)
}

func TestEventInMemory_Delete_NotExist(t *testing.T) {
	repo, ctx := setupEventRepo()

	// 刪除不存在項目應該不報錯
	err := repo.Delete(ctx, "nothing")
	require.NoError(t, err)
}
