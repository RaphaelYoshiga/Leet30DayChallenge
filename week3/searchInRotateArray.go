package main;

func search(nums []int, target int) int {

	left:= 0;
	right:= len(nums) -1;

	for left <= right{
		mid := (left + right)/2;

		midElement := nums[mid];
		curLeftEle := nums[left];
		curRightEle := nums[right];

		if midElement == target{
			return mid;
		}

		// Mid Element in the left section
		if midElement > curRightEle {
			if target <= curLeftEle && target < midElement {
				right = mid -1;
			}else{
				left = mid + 1;
			}
		}else if midElement < curLeftEle{ // Element in the right section
			if target <= curRightEle && target > midElement{
				left = mid+1;
			}else{
				right = mid -1;
			}
		}else{
			if target > midElement {
				left = mid +1;
			} else{
				right = mid -1;
			}
		}
	}
	return -1;
}