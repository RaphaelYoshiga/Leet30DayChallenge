package main

import "testing"

func TestLruCache(t *testing.T) {
	obj := NewLruCache(2);
	
	obj.Put(1, 1);
	obj.Put(2, 2);
	
	obj.Get(1);
	obj.Get(2);
	
	obj.Put(3, 2);

	result := obj.Get(1);
	if result != -1{
		t.Errorf("Cache %d key should have been invalidated but returned: %d", 1, result);
	}
}


func TestLruCacheAnotherScenario2(t *testing.T) {
	obj := NewLruCache(2);
	
	obj.Get(2);
	obj.Put(2, 6);

	obj.Get(1);
	obj.Put(1, 5);
	obj.Put(1, 2);

	getAndAssert(t, &obj, 1, 2);
	getAndAssert(t, &obj, 2, 6);
}


func TestLruCacheAnotherScenario3(t *testing.T) {
	obj := NewLruCache(2);
	
	obj.Put(1, 1);
	obj.Put(2, 2);
	
	obj.Get(1);
	obj.Put(3, 3);

	obj.Get(2);
	obj.Put(4, 4);

	getAndAssert(t, &obj, 1, -1);
	getAndAssert(t, &obj, 3, 3);
	getAndAssert(t, &obj, 4, 4);
}

func TestLruCacheAnotherScenario(t *testing.T) {
	obj := NewLruCache(2);
	
	obj.Put(2, 1);
	obj.Put(1, 1);
	obj.Put(2, 3);
	obj.Put(4, 1);
	
	getAndAssert(t, &obj, 1, -1);
	getAndAssert(t, &obj, 2, 3);
	getAndAssert(t, &obj, 4, 1);
}

func getAndAssert(t *testing.T, cache *LRUCache, key int, expected int){
	result := cache.Get(key);
	if result != expected{
		t.Errorf("Cache %d key should be %d returned: %d", key, expected, result);
	}
}