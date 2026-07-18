package cache

import (
    "testing"
    "time"
)

func TestSetAndGetDefaultCache(t *testing.T) {
    c := NewCache(LRU)

    c.Set("name", "jp", 10)

    value, ok := c.Get("name")
    if !ok {
        t.Fatal("expected key to exist")
    }

    if value != "jp" {
        t.Fatalf("expected value 'jp', got '%s'", value)
    }
}

func TestDeleteDefaultCache(t *testing.T) {
    c := NewCache(LRU)

    c.Set("name", "jp", 10)
    c.Del("name")

    _, ok := c.Get("name")
    if ok {
        t.Fatal("expected key to be deleted")
    }
}

func TestTTLExpirationDefaultCache(t *testing.T) {
    c := NewCache(LRU)

    c.Set("temp", "data", 1)

    time.Sleep(2 * time.Second)

    _, ok := c.Get("temp")
    if ok {
        t.Fatal("expected key to expire")
    }
}

func TestOverwriteKeyDefaultCache(t *testing.T) {
    c := NewCache(LRU)

    c.Set("name", "jp", 10)
    c.Set("name", "john", 10)

    value, ok := c.Get("name")
    if !ok {
        t.Fatal("expected key to exist")
    }

    if value != "john" {
        t.Fatalf("expected updated value 'john', got '%s'", value)
    }
}

func TestLRUEviction(t *testing.T) {
    c := NewCache(LRU)

    c.Set("a", "1", 10)
    c.Set("b", "2", 10)
    c.Set("c", "3", 10)

    _, ok := c.Get("a")
    if !ok {
        t.Fatal("expected key a to exist")
    }

    c.Set("d", "4", 10)

    if _, ok := c.Get("b"); ok {
        t.Fatal("expected key b to be evicted by LRU")
    }

    for _, key := range []string{"a", "c", "d"} {
        if _, ok := c.Get(key); !ok {
            t.Fatalf("expected key %s to still exist", key)
        }
    }
}

func TestLFUEviction(t *testing.T) {
    c := NewCache(LFU)

    c.Set("a", "1", 10)
    c.Set("b", "2", 10)
    c.Set("c", "3", 10)

    _, ok := c.Get("a")
    if !ok {
        t.Fatal("expected key a to exist")
    }

    _, ok = c.Get("a")
    if !ok {
        t.Fatal("expected key a to exist on second read")
    }

    c.Set("d", "4", 10)

    if _, ok := c.Get("b"); ok {
        t.Fatal("expected key b to be evicted by LFU")
    }

    for _, key := range []string{"a", "c", "d"} {
        if _, ok := c.Get(key); !ok {
            t.Fatalf("expected key %s to still exist", key)
        }
    }
}

func TestNewCacheWithNilEvictionFallsBackToLRU(t *testing.T) {
    c := NewCacheWithEviction(nil)

    c.Set("name", "jp", 10)

    value, ok := c.Get("name")
    if !ok {
        t.Fatal("expected key to exist")
    }

    if value != "jp" {
        t.Fatalf("expected value 'jp', got '%s'", value)
    }
}