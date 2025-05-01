package repository

import (
	"errors"
	"testing"

	"backend/internal/entity"
)

func TestNewCounterInMemory_WithNilMap(t *testing.T) {
	cm := NewCounterInMemory(nil)
	if cm == nil || cm.counters == nil {
		t.Fatalf("NewCounterInMemory returned invalid instance")
	}
}

func TestNewCounterInMemory_WithInjectedMap(t *testing.T) {
	mockCounter := entity.NewCounter(10, nil)
	cm := NewCounterInMemory(map[string]*entity.Counter{
		"preloaded": mockCounter,
	})

	got, err := cm.Get("preloaded")
	if err != nil {
		t.Errorf("expected to get preloaded counter, got error: %v", err)
	}
	if got != mockCounter {
		t.Errorf("expected same pointer for preloaded counter")
	}
}

func TestCounterInMemory_NewAndGet(t *testing.T) {
	cm := NewCounterInMemory(nil)

	counter, err := cm.New("alpha", 5, nil)
	if err != nil {
		t.Fatalf("unexpected error on New: %v", err)
	}

	got, err := cm.Get("alpha")
	if err != nil {
		t.Errorf("unexpected error on Get: %v", err)
	}
	if got != counter {
		t.Errorf("expected same pointer returned from Get")
	}
}

func TestCounterInMemory_New_Duplicate(t *testing.T) {
	cm := NewCounterInMemory(nil)
	_, _ = cm.New("dup", 1, nil)

	_, err := cm.New("dup", 1, nil)
	if !errors.Is(err, ErrCounterExist) {
		t.Errorf("expected ErrCounterExist, got %v", err)
	}
}

func TestCounterInMemory_Get_NotFound(t *testing.T) {
	cm := NewCounterInMemory(nil)

	_, err := cm.Get("ghost")
	if !errors.Is(err, ErrCannotGetCounter) {
		t.Errorf("expected ErrCannotGetCounter, got %v", err)
	}
}

func TestCounterInMemory_Delete_Exist(t *testing.T) {
	cm := NewCounterInMemory(nil)
	_, _ = cm.New("temp", 1, nil)

	err := cm.Delete("temp")
	if err != nil {
		t.Errorf("unexpected error on Delete: %v", err)
	}

	_, err = cm.Get("temp")
	if !errors.Is(err, ErrCannotGetCounter) {
		t.Errorf("expected counter to be deleted, but still found")
	}
}

func TestCounterInMemory_Delete_NotExist(t *testing.T) {
	cm := NewCounterInMemory(nil)

	err := cm.Delete("nothing")
	if err != nil {
		t.Errorf("Delete on nonexistent counter should not error")
	}
}

func TestCounterInMemory_List_Sorted(t *testing.T) {
	cm := NewCounterInMemory(nil)
	cm.New("b", 1, nil)
	cm.New("a", 1, nil)
	cm.New("c", 1, nil)

	list := cm.List()
	if len(list) != 3 {
		t.Errorf("expected 3 counters, got %d", len(list))
	}

	// 確認排序順序 a, b, c
	expectedOrder := []string{"a", "b", "c"}
	for i, name := range expectedOrder {
		counter, _ := cm.Get(name)
		if list[i] != counter {
			t.Errorf("expected %s at position %d", name, i)
		}
	}
}
