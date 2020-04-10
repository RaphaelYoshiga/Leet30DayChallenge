package main

func maxSubArray(nums []int) int {
	max:= nums[0];
	sum := 0;

	for _, num := range nums {
		if num < 0 {
			if num + sum > 0 {
				sum += num;
			} else {
				sum = 0;
				max = takeMax(max, num);
			}
		} else{
			sum += num;
			max = takeMax(max, sum);
		}	
	}
	return max;
}

func takeMax(a int, b int) int{
	if a > b {
		return a;
	}
	return b;
}

