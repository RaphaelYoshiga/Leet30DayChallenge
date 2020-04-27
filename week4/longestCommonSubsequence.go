package main;

func longestCommonSubsequence(a string, b string) int {
	var dp = make([][]int, len(a) + 1)
	for i := range dp {
		dp[i] = make([]int, len(b) + 1)
	}

	a = " " + a;
	b = " " + b;
	for i:= 1; i < len(dp); i++{
		for j:= 1; j<len(dp[0]); j++{
			if a[i] == b[j]{
				dp[i][j] = 1 + dp[i -1][j-1];
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0]) -1]
}

