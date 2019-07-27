package search

// Max returns the larger of two integers.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// LCSImpl computes the Longest Common Subsequence problem of two strings.
func LCSImpl(pattern string, target string) (hit bool, matchCount int) {
	hit = false
	lenPattern, lenTarget := len(pattern), len(target)
	dp := make([][]int, lenPattern+1)
	for i := 0; i < lenPattern+1; i++ {
		dp[i] = make([]int, lenTarget+1)
	}

	for i := 1; i <= lenPattern; i++ {
		for j := 1; j <= lenTarget; j++ {
			if pattern[i-1] == target[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	matchCount = dp[lenPattern][lenTarget]
	if matchCount == lenPattern {
		hit = true
	}
	return
}

// LCS search from the end to the beginning.
// This fits our behavior better on path searching.
// Note that one could also perform string reversing followed by normal LCS
// but this function takes slightly better care of performance.
func reverseLCSImpl(pattern string, target string) (hit bool, matchCount int) {
	hit = false
	lenPattern, lenTarget := len(pattern), len(target)
	dp := make([][]int, lenPattern+1)
	for i := 0; i < lenPattern+1; i++ {
		dp[i] = make([]int, lenTarget+1)
	}

	for i := lenPattern - 1; i >= 0; i-- {
		for j := lenTarget - 1; j >= 0; j-- {
			if pattern[i] == target[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	matchCount = dp[0][0]
	if matchCount == lenPattern {
		hit = true
	}
	return
}
