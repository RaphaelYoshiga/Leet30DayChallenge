package main

func moveZeroes(nums []int) {
	l := len(nums); 

	c:=0;
	for i := 0; i < l; i++ {
		a:= nums[i];

		if a == 0 {
			c++;			
		} else { 
			nums[i - c] = a;
		}
	}
	
	for c > 0 {
		nums[l - c] = 0;

		c--;
	}
}