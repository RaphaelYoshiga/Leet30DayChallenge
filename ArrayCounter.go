package main;

func countElements(arr []int) int {

	m := make(map[int]int);

	for _, num := range arr{
		m[num]++;
	}

	total := 0;
	for key, numberCount := range m{
		if m[key + 1] > 0 {
			total += numberCount;
		}

	}
	return total;
}