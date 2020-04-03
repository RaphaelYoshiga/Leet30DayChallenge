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
	sum := 0
	for n > 0 {
		sum+= (n % 10) * (n % 10);
		n = n / 10;
	}
    return sum;
}

