package main;

func findMaxLength(nums []int) int {
	if len(nums) < 2{
		return 0;
	}

	m := make(map[int]int);

	maxLenght:= 0;
	sum := 0;
	for i, n := range nums{

		if n == 0{
			sum--;
		}else{
			sum++;
		}
		
		switch {
			case sum == 0:
				maxLenght = max(maxLenght,i+1)
				break;
			case m[sum] > 0:
				maxLenght = max(maxLenght,  i-m[sum] +1 ) 
				break;
			default:
				m[sum] = i+1;
		}
	}

	return maxLenght;
}


func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func main(){}