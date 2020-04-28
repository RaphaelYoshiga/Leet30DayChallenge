package main;

type UniqueNode struct {
	val *int
	prev *UniqueNode
	next *UniqueNode
}
type FirstUnique struct {
	nodeMap map[int]*UniqueNode
	head *UniqueNode
	tail *UniqueNode
}


func Constructor(nums []int) FirstUnique {
	head:= &UniqueNode {  }
	tail:= &UniqueNode { prev: head	};
	head.next = tail;

	firstUnique := FirstUnique { head: head, tail: tail, nodeMap: make(map[int]*UniqueNode)};

	for _, n := range nums {
		firstUnique.Add(n)
	}

	return firstUnique;
}


func (this *FirstUnique) ShowFirstUnique() int {
	val:=this.head.next.val;
	if val == nil{
		return -1
	}
	return *val;
}


func (this *FirstUnique) Add(value int)  {

	foundNode, ok := this.nodeMap[value];

	if ok {
		tryRemove(this, foundNode)
	}else{
		node := &UniqueNode { val: &value, }
		addToQueue(this, node);
		this.nodeMap[value] = node;
	}

}

func tryRemove(firstUnique *FirstUnique, node *UniqueNode)  {

	if node.prev != nil{

		prevNode := node.prev;
		nextNode := node.next;

		prevNode.next = nextNode;
		nextNode.prev = prevNode;

		node.next =nil;
		node.prev =nil;
	}
}


func addToQueue(firstUnique *FirstUnique, node *UniqueNode)  {
	tailPrev := firstUnique.tail.prev;

	tailPrev.next = node;
	firstUnique.tail.prev = node;
	
	node.prev = tailPrev;
	node.next = firstUnique.tail;
}

func main(){
	obj := Constructor([]int { 2, 3, 5} );
    
	a := obj.ShowFirstUnique();;
	a = a -1+1;
    obj.Add(5);            // the queue is now [2,3,5,5]
    obj.ShowFirstUnique();;
    obj.Add(2);            // the queue is now [2,3,5,5,2]
    obj.ShowFirstUnique();;
    obj.Add(3);            // the queue is now [2,3,5,5,2,3]
    obj.ShowFirstUnique();;
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */