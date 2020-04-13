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
		
		if sum == 0{
			maxLenght = i + 1 
		}

		if _, ok := m[sum]; ok {
            if maxLenght < i - m[sum] {
				maxLenght = i - m[sum] 
			}
		}else{
			m[sum] = i;
		}
	}

	return maxLenght;
}

func min(a int, b int) int{
	if a < b {
		return a;
	}
	return b;
}

func main(){
	findMaxLength( []int {0,1,0,0,1});
}