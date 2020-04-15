package main

func productExceptSelf(nums []int) []int {

	r := []int {};

	for i, _ := range nums {
		r = append(r, getProductSubset(i, nums));
	}

	return r;
}

func getProductSubset(i int, nums []int) int{
	if i == 0 {
		return product(nums[1:len(nums)]);
	}

	ending:= nums[i+1:len(nums)];
	beginning := nums[0:i];

	product := 1;

	for _, l := range ending{
		product *= l;
	}
	for _, l := range beginning{
		product *= l;
	}
	return product;
}

func product(nums []int) int{
	product := 1;

	for _, l := range nums{
		product *= l;
	}
	return product;
}

func main(){

	productExceptSelf([]int {1,0});
}