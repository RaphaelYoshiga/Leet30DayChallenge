package main;

type LRUCache struct {
	keys map[int]*Node
	capacity int
	head *Node
	tail *Node
}

type Node struct{
	key int
	val int

	prev *Node
	next *Node
}

func NewLruCache(capacity int) LRUCache {
	head:= Node {}
	tail := Node {};

	head.next = &tail;
	tail.prev = &head;

	return LRUCache{ keys: make(map[int]*Node), capacity: capacity, head: &head, tail: &tail };
}

func (this *LRUCache) Get(key int) int {
	
	node, ok := this.keys[key]
	if ok {
		remove(this, node);
		add(this, node);		
		return this.keys[key].val;
	}
	return -1;
}

func add (cache *LRUCache, node *Node){
	headNext:= cache.head.next;

	cache.head.next = node;
	headNext.prev = node;

	node.prev = cache.head;
	node.next = headNext;
}

func remove (cache *LRUCache, node *Node){

	nextNode := node.next;
	prevNode := node.prev;

	nextNode.prev = prevNode;
	prevNode.next = nextNode;
}

func (this *LRUCache) Put(key int, value int)  {

	_, ok := this.keys[key]

	if ok {
		node:= this.keys[key];
		node.val = value;

		remove(this, node);
		add(this, node)
	} else {
		if len(this.keys) == this.capacity {
			nodeToRemove :=this.tail.prev;
			delete(this.keys, nodeToRemove.key)
			remove(this, nodeToRemove);
		}

		node:= Node { key: key, val: value };;
		add(this, &node)
		this.keys[key] = &node;
	}
}

// Previous implementation 
//func sortUsedQueue(cache *LRUCache, key int){
// 	found:= false;
// 	for i, n := range cache.used{
// 		if n == key{
// 			cache.used = append(cache.used[:i], cache.used[i+1:]...)
// 			cache.used = append(cache.used, key)
// 			found = true;
// 			break;
// 		}
// 	}

// 	if !found{
// 		cache.used = append(cache.used, key)
// 	}
// }

// func (this *LRUCache) Get(key int) int {
	
// 	_, ok := this.keys[key]
// 	if ok {
// 		sortUsedQueue(this, key);
// 		return this.keys[key].val;
// 	}
// 	return -1;
// }

// func (this *LRUCache) Put(key int, value int)  {
// 	sortUsedQueue(this, key);

// 	_, ok := this.keys[key]
// 	if len(this.keys) == this.capacity && !ok{

// 		i := 0;
// 		for i < len(this.used) -1{
// 			keyToDelete := this.used[i];
// 			_, inCache := this.keys[keyToDelete]
// 			if inCache{
// 				this.used = append(this.used[:i], this.used[i+1:]...)
// 				delete(this.keys, keyToDelete);
// 				break;
// 			}
// 			i++;
// 		}

// 	}
		
// 	this.keys[key] = value;
// }
