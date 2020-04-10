package main

func isHappy(n int) bool {
	turtle := n;
	hare := n;

	for true {
		turtle = digitsSum(turtle)
		hare = digitsSum(digitsSum(hare));

		if hare == 1 {
			return true;
		} else if turtle == hare {
			break;
		}
    }

	return false;
}

func digitsSum(n int) int {
	sum := 0
	for n > 0 {
		sum+= (n % 10) * (n % 10);
		n = n / 10;
	}
    return sum;
}

