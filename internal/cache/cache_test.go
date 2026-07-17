package cache

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	cache := NewCache()

	cache.Set("name", "jp", 10)

	ok, value := cache.Get("name")

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "jp" {
		t.Fatalf("expected value 'jp', got '%s'", value)
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache()

	cache.Set("name", "jp", 10)

	cache.Del("name")

	ok, _ := cache.Get("name")

	if ok {
		t.Fatal("expected key to be deleted")
	}
}

func TestTTLExpiration(t *testing.T) {
	cache := NewCache()

	cache.Set("temp", "data", 1)

	time.Sleep(2 * time.Second)

	ok, _ := cache.Get("temp")

	if ok {
		t.Fatal("expected key to expire")
	}
}

func TestOverwriteKey(t *testing.T) {
	cache := NewCache()

	cache.Set("name", "jp", 10)
	cache.Set("name", "john", 10)

	ok, value := cache.Get("name")

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "john" {
		t.Fatalf("expected updated value 'john', got '%s'", value)
	}
}