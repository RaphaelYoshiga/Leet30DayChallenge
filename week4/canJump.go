package main;

func canJump(nums []int) bool {
	target := len(nums) -1;

	lastGoodIndex := target;
	for i:= target; i>= 0; i--{
		if i + nums[i] >= lastGoodIndex {
			lastGoodIndex = i;
		}
	}
	
	return lastGoodIndex == 0;
}