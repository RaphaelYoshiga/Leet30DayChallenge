package main

func moveZeroes(nums []int) {
	maxIterations := len(nums) - 1; 
	for i := 0; i < maxIterations; i++ {
		a:= nums[i];

		if a == 0 {
			bi := i+ 1;
			
			b:= nums[bi];

			for b == 0 {
				if bi == len(nums) -1{
					break;
				}
				
				bi++;
				b = nums[bi]
			}

			nums[i] = b;
			nums[bi] = a;
		}
		
	}
}