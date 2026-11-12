package cache

import (
	"sync"
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {
	c := New()

	c.Set("node_status", "online", 1*time.Second)

	val, found := c.Get("node_status")
	if !found {
		t.Fatalf("expected key to exist")
	}

	if val != "online" {
		t.Errorf("expected 'online', got %v", val)
	}
}

func TestCacheExpiration(t *testing.T) {
	c := New()

	// Set TTL for 50 milliseconds
	c.Set("temp_key", "data", 50*time.Millisecond)

	// Wait 70 milliseconds
	time.Sleep(70 * time.Millisecond)

	_, found := c.Get("temp_key")
	if found {
		t.Errorf("expected key to be expired, but it was found")
	}
}

// TestConcurrentAccess simulates 100 goroutines reading and writing simultaneously
func TestConcurrentAccess(t *testing.T) {
	c := New()
	var wg sync.WaitGroup

	// 50 concurrent writers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Set("key", id, 100*time.Millisecond)
		}(i)
	}

	// 50 concurrent readers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("key")
		}()
	}

	wg.Wait()
}