package main;

type LRUCache struct {
	keys map[int]int
	capacity int
	used []int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{ keys: make(map[int]int), capacity: capacity };
}

func sortUsedQueue(cache *LRUCache, key int){
	found:= false;
	for i, n := range cache.used{
		if n == key{
			cache.used = append(cache.used[:i], cache.used[i+1:]...)
			cache.used = append(cache.used, key)
			found = true;
			break;
		}
	}

	if !found{
		cache.used = append(cache.used, key)
	}
}

func (this *LRUCache) Get(key int) int {
	
	_, ok := this.keys[key]
	if ok {
		sortUsedQueue(this, key);
		return this.keys[key];
	}
	return -1;
}

func (this *LRUCache) Put(key int, value int)  {
	sortUsedQueue(this, key);

	_, ok := this.keys[key]
	if len(this.keys) == this.capacity && !ok{

		i := 0;
		for i < len(this.used) -1{
			keyToDelete := this.used[i];
			_, inCache := this.keys[keyToDelete]
			if inCache{
				this.used = append(this.used[:i], this.used[i+1:]...)
				delete(this.keys, keyToDelete);
				break;
			}
			i++;
		}

	}
		
	this.keys[key] = value;
}


func main(){

	obj := Constructor(2);
	
	obj.Put(1, 1);
	obj.Put(2, 2);
	
	obj.Get(1);
	obj.Put(3, 3);

	obj.Get(2);
	obj.Put(4, 4);

	res := obj.Get(1);
	res++;
		
}