package entity

import (
	"testing"
	"time"
)

// stub callback，記錄 callback 是否被呼叫
type callbackRecorder struct {
	called int
	last   *Counter
}

func (r *callbackRecorder) callback(c *Counter) {
	r.called++
	r.last = c
}

func TestNewCounter(t *testing.T) {
	var recorder callbackRecorder
	c := NewCounter(5, recorder.callback)

	if c.Count() != 5 {
		t.Errorf("expectedinitial count to be 5, got %d", c.Count())
	}
	if c.Counting() {
		t.Errorf("expected counting to be false at init")
	}
}

func TestSetInit(t *testing.T) {
	var recorder callbackRecorder
	c := NewCounter(3, recorder.callback)

	c.SetInit(10)
	if c.Count() != 10 {
		t.Errorf("SetInit failed, expected 10, got %d", c.Count())
	}
	if c.init != 10 {
		t.Errorf("SetInit failed, init value not updated")
	}
	if recorder.called != 1 {
		t.Errorf("expected callback to be called once, got %d", recorder.called)
	}
}

func TestCounterReset(t *testing.T) {
	var recorder callbackRecorder
	c := NewCounter(3, recorder.callback)
	c.count = 1

	c.Reset()
	if c.Count() != 3 {
		t.Errorf("Reset failed, expected count to be 3, got %d", c.Count())
	}
	if recorder.called != 2 { // Stop + Reset 都觸發 callback
		t.Errorf("expected callback to be called twice, got %d", recorder.called)
	}
}

func TestStopWhenNotCounting(t *testing.T) {
	var recorder callbackRecorder
	c := NewCounter(2, recorder.callback)
	c.Stop()
	if recorder.called != 1 {
		t.Errorf("Stop should still trigger callback, got %d", recorder.called)
	}
}

func TestStartAndStop(t *testing.T) {
	var recorder callbackRecorder
	c := NewCounter(1, recorder.callback)
	c.Start()
	time.Sleep(1100 * time.Millisecond) // 等待一秒讓 tick 發生
	if c.Counting() {
		t.Errorf("expected counting to be false after countdown")
	}
	if c.Count() != 0 {
		t.Errorf("expected counter to reach 0, got %d", c.Count())
	}
	if recorder.called == 0 {
		t.Errorf("expected callback to be triggered")
	}
}

func TestNewCounterWithNilCallback(t *testing.T) {
	// 不傳 callback，應該不 panic
	c := NewCounter(3, nil)

	if c.Count() != 3 {
		t.Errorf("expected initial count to be 3, got %d", c.Count())
	}

	// 確保呼叫包含 callback 的操作也不會出錯（SetInit 有觸發 callback）
	c.SetInit(5)
	if c.Count() != 5 {
		t.Errorf("expected count to be 5 after SetInit, got %d", c.Count())
	}
}

func TestSetInitWhenCounting(t *testing.T) {
	c := NewCounter(5, nil)
	c.Start()
	time.Sleep(1100 * time.Millisecond) // 等待一秒讓 tick 發生
	c.SetInit(10)
	if c.Count() != 10 {
		t.Errorf("SetInit failed, expected 10, got %d", c.Count())
	}
	if c.Counting() {
		t.Errorf("expected counting to be false after SetInit")
	}
}

func TestStopWhenCounting(t *testing.T) {
	c := NewCounter(5, nil)
	c.Start()
	time.Sleep(1100 * time.Millisecond) // 等待一秒讓 tick 發生
	c.Stop()
	if c.Counting() {
		t.Errorf("expected counting to be false after Stop")
	}
}
