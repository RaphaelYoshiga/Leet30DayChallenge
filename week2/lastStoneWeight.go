package main;

import "sort";

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0;
	}

	if len(stones) == 1 {
		return 1;
	}

	for len(stones) > 1{

		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})

		fightResult := abs(stones[0] - stones[1]);

		if len(stones) == 2 {
			return fightResult;
		}	

		if fightResult > 0{
			stones = stones[1:len(stones)]
			stones[0] = fightResult;
		}else{
			stones = stones[2:len(stones)]
		}
	}

	return stones[0];
}

func abs(n int) int{
	if n> 0 {
		return n;
	}
	return -n;
}
func main(){
	lastStoneWeight( []int {3,7,2});
}
