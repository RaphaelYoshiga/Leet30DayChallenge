package main;

func rangeBitwiseAnd(m int, n int) int {

	bitwise:= m;
	m++;

	for m <= n{
		bitwise &= m
		m++;
	}
		
	return bitwise;
}