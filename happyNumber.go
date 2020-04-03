package main

var h map[int]bool = make(map[int]bool) 

func isHappy(n int) bool {
	
	current := splitAndSum(n);
	if h[current] && current != 1 {
		h = make(map[int]bool);
		return false;
	}
	h[current] = true    

	if current == 1	{
		h = make(map[int]bool);
		return true;
	}
	return isHappy(current);
}


func splitAndSum(n int) int {
	x := []int { }
	for n > 0 {
		x = append(x, n % 10)
		n = n / 10;
	}

	sum := 0
	for _, num := range x {
		sum += num * num;
	}
    return sum;
}

