package main

func subarraySum(nums []int, k int) int {
	m := make(map[int]int);

	m[0] = 1;
	sum := 0;
	result := 0;

	for _, n:= range nums{
		sum += n;
		if m[sum -k] > 0{
			result+= m[sum -k];
		}

		m[sum] += 1;
	}

	return result;
}

