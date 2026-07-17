package cache

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	cache := NewCache()

	cache.Set("user1", "jyoti")

	found, val := cache.Get("user1")

	if !found {
		t.Fatal("expected key to exist")
	}

	if val != "jyoti" {
		t.Fatalf("expected jyoti, got %s", val)
	}
}

func TestMissingKey(t *testing.T) {
	cache := NewCache()

	found, _ := cache.Get("missing")

	if found {
		t.Fatal("expected key to not exist")
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache()

	cache.Set("user1", "jyoti")
	cache.Del("user1")

	found, _ := cache.Get("user1")

	if found {
		t.Fatal("expected key to be deleted")
	}
}

func TestOverwrite(t *testing.T) {
	cache := NewCache()

	cache.Set("user1", "old")
	cache.Set("user1", "new")

	found, val := cache.Get("user1")

	if !found {
		t.Fatal("expected key to exist")
	}

	if val != "new" {
		t.Fatalf("expected new, got %s", val)
	}
}