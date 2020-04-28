package main;

import "testing";

func TestSimpleUniqueQueue(t *testing.T) {
	obj := Constructor([]int { 2, 3, 5} );
    
    AssertFirstUnique(&obj, 2, t); // return 2
    obj.Add(5);            // the queue is now [2,3,5,5]
    AssertFirstUnique(&obj, 2, t); // return 2
    obj.Add(2);            // the queue is now [2,3,5,5,2]
    AssertFirstUnique(&obj, 3, t);
    obj.Add(3);            // the queue is now [2,3,5,5,2,3]
    AssertFirstUnique(&obj, -1, t);
}

func AssertFirstUnique(obj *FirstUnique, expected int, t *testing.T){
    result :=  obj.ShowFirstUnique();;
	if result != result{
		t.Errorf("First unique item is %d, but should be %d", result, expected);
	}

}
