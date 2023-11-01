package cache

import (
	"testing"
	"time"
)

func TestCache_SetAndGetWithStruct(t *testing.T) {
	type Foo struct {
		Bar string
		Baz int
	}

	cache := NewCache[string, Foo]()

	cache.Set("foo", Foo{"bar", 42}, time.Minute)

	val, ok := cache.Get("foo")
	if !ok {
		t.Errorf("Expected to get value for key 'foo', but got none")
	}
	if val.Bar != "bar" {
		t.Errorf("Expected value for key 'foo' to be 'bar', but got %s", val.Bar)
	}
	if val.Baz != 42 {
		t.Errorf("Expected value for key 'foo' to be 42, but got %d", val.Baz)
	}

	_, ok = cache.Get("bar")
	if ok {
		t.Errorf("Expected to get no value for key 'bar', but got one")
	}
}

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache[string, int]()

	cache.Set("foo", 42, time.Minute)

	val, ok := cache.Get("foo")
	if !ok {
		t.Errorf("Expected to get value for key 'foo', but got none")
	}
	if val != 42 {
		t.Errorf("Expected value for key 'foo' to be 42, but got %d", val)
	}

	_, ok = cache.Get("bar")
	if ok {
		t.Errorf("Expected to get no value for key 'bar', but got one")
	}
}

func TestCache_Exist(t *testing.T) {
	cache := NewCache[string, int]()

	cache.Set("foo", 42, time.Minute)

	exists := cache.Exist("foo")
	if !exists {
		t.Errorf("Expected key 'foo' to exist, but it doesn't")
	}

	exists = cache.Exist("bar")
	if exists {
		t.Errorf("Expected key 'bar' to not exist, but it does")
	}
}
