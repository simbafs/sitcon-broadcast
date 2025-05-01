package repository

// 我也不知道這個東西幹麻寫測試，但反正用 AI 寫很快，就放在這裡了，增加 cover rate

import (
	"context"
	"testing"

	"backend/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setupTestRepo(t *testing.T) (*EventEnt, context.Context) {
	client := enttest.Open(t, "sqlite3", "file:event?mode=memory&cache=shared&_fk=1")
	return NewEventEnt(client), context.Background()
}

func TestEventEnt_CreateAndGet(t *testing.T) {
	repo, ctx := setupTestRepo(t)
	defer repo.client.Close()

	ev, err := repo.Create(ctx, "test", "https://url", "console.log('hi')")
	require.NoError(t, err)
	require.Equal(t, "test", ev.Name())
	require.Equal(t, "https://url", ev.URL())

	// Test Get
	got, err := repo.Get(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, ev.Name(), got.Name())
}

func TestEventEnt_List(t *testing.T) {
	repo, ctx := setupTestRepo(t)
	defer repo.client.Close()

	repo.Create(ctx, "b", "urlB", "b()")
	repo.Create(ctx, "a", "urlA", "a()")
	repo.Create(ctx, "c", "urlC", "c()")

	list, err := repo.List(ctx)
	require.NoError(t, err)
	require.Len(t, list, 3)
	require.Equal(t, "a", list[0].Name()) // List 是按照 name 排序
}

func TestEventEnt_Update(t *testing.T) {
	repo, ctx := setupTestRepo(t)
	defer repo.client.Close()

	repo.Create(ctx, "event", "old-url", "old-script")

	err := repo.Update(ctx, "event", "new-url", "new-script")
	require.NoError(t, err)

	got, _ := repo.Get(ctx, "event")
	require.Equal(t, "new-url", got.URL())
	require.Equal(t, "new-script", got.Script())
}

func TestEventEnt_Update_NotFound(t *testing.T) {
	repo, ctx := setupTestRepo(t)
	defer repo.client.Close()

	err := repo.Update(ctx, "ghost", "x", "y")
	require.Equal(t, ErrEventNotFound, err)
}

func TestEventEnt_Delete(t *testing.T) {
	repo, ctx := setupTestRepo(t)
	defer repo.client.Close()

	repo.Create(ctx, "to-delete", "url", "script")

	err := repo.Delete(ctx, "to-delete")
	require.NoError(t, err)

	_, err = repo.Get(ctx, "to-delete")
	require.Error(t, err)
}
