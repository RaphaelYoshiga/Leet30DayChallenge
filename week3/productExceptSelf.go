package main

func productExceptSelf(nums []int) []int {

	r := []int { 1 };
	l := len(nums);
	for i := 1; i < l; i++ {
		product:= nums[i-1] * r[i-1];
		r = append(r, product);
	}
	
	secondProduct := 1;
	for i := l -2; i >= 0; i-- {
		secondProduct = nums[i+1] * secondProduct;
		r[i] *= secondProduct; 
	}

	return r;
}
