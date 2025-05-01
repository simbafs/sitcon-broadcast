package entity

import (
	"testing"
	"time"
)

// 測試初始化後 Get() 回傳指定時間
func TestNewNow(t *testing.T) {
	n := NewNow(12345)
	if got := n.Get(); got != 12345 {
		t.Errorf("expected Get() to return 12345, got %d", got)
	}
}

// 測試 Set()
func TestSet(t *testing.T) {
	n := NewNow(0)
	n.Set(54321)
	if got := n.Get(); got != 54321 {
		t.Errorf("expected Get() after Set() to return 54321, got %d", got)
	}
}

// 測試 Reset()
func TestReset(t *testing.T) {
	n := NewNow(99999)
	n.Reset()
	got := n.Get()
	now := time.Now().Unix()

	if got < now-1 || got > now+1 {
		t.Errorf("expected Get() after Reset() to return near current time, got %d, want around %d", got, now)
	}
}

// 測試 Get() 在 now == 0 時的 fallback 行為
func TestGetWithZero(t *testing.T) {
	n := NewNow(0)
	got := n.Get()
	now := time.Now().Unix()

	if got < now-1 || got > now+1 {
		t.Errorf("expected Get() with zero value to return current time, got %d, want around %d", got, now)
	}
}

// ✅ 測試重複 Set、Reset 並觀察行為是否穩定（多一層 safety）
func TestSetResetSequence(t *testing.T) {
	n := NewNow(0)
	n.Set(100)
	if got := n.Get(); got != 100 {
		t.Errorf("after Set(100), expected 100, got %d", got)
	}

	n.Reset()
	now1 := n.Get()
	n.Set(200)
	if got := n.Get(); got != 200 {
		t.Errorf("after Set(200), expected 200, got %d", got)
	}

	n.Reset()
	now2 := n.Get()

	if now2 < now1-1 || now2 > now1+5 {
		t.Errorf("expected time after second reset to be close to current, got %d", now2)
	}
}
