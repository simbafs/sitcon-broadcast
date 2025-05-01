package entity

import "testing"

func TestNewEvent(t *testing.T) {
	e := NewEvent("event1", "http://example.com", "console.log('hi')")

	if e.Name() != "event1" {
		t.Errorf("expected name to be 'event1', got %s", e.Name())
	}
	if e.URL() != "http://example.com" {
		t.Errorf("expected url to be 'http://example.com', got %s", e.URL())
	}
	if e.Script() != "console.log('hi')" {
		t.Errorf("expected script to be 'console.log('hi')', got %s", e.Script())
	}
}

func TestSetURL(t *testing.T) {
	e := NewEvent("test", "old.com", "script")
	e.SetURL("new.com")
	if e.URL() != "new.com" {
		t.Errorf("expected URL to be 'new.com', got %s", e.URL())
	}
}

func TestSetScript(t *testing.T) {
	e := NewEvent("test", "url", "old-script")
	e.SetScript("new-script")
	if e.Script() != "new-script" {
		t.Errorf("expected script to be 'new-script', got %s", e.Script())
	}
}
