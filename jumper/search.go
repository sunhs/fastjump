package jumper

import (
	"os"
	"path/filepath"
	"strings"
)

type tup struct {
	index int
	path  string
}

func MatchDispatcher(patterns []string, candidates []string, count int, sep string) (rst []tup) {
	cwd := os.Getenv("PWD")

	for idx, candidate := range candidates {
		if candidate == cwd {
			continue
		} else if candidateAbs, _ := filepath.Abs(candidate); candidateAbs == cwd {
			continue
		} else if info, err := os.Stat(candidate); err != nil || !info.IsDir() {
			continue
		}

		candParts := strings.Split(candidate, sep)
		candPartsLower := make([]string, len(candParts))
		for i, candPart := range candParts {
			candPartsLower[i] = strings.ToLower(candPart)
		}

		patternsLower := make([]string, len(patterns))
		for i, pattern := range patterns {
			patternsLower[i] = strings.ToLower(pattern)
		}

		if partMatch(patternsLower, candPartsLower) {
			rst = append(rst, tup{idx, candidate})
		}

		if len(rst) == count {
			return
		}
	}

	if len(patterns) == 1 {
		for idx, candidate := range candidates {
			if candidate == cwd {
				continue
			} else if info, err := os.Stat(candidate); err != nil || !info.IsDir() {
				continue
			}

			if wholeMatch(strings.ToLower(patterns[0]), strings.ToLower(candidate)) {
				rst = append(rst, tup{idx, candidate})
			}

			if len(rst) == count {
				return
			}
		}
	}

	return
}

func wholeMatch(pattern string, candidate string) bool {
	hit, _ := reverseLCSImpl(pattern, candidate)
	return hit
}

func partMatch(patterns []string, candidateParts []string) bool {
	revPats := make([]string, len(patterns))
	for i := len(patterns) - 1; i >= 0; i-- {
		revPats[len(patterns)-1-i] = patterns[i]
	}

	revCandParts := make([]string, len(candidateParts))
	for i := len(candidateParts) - 1; i >= 0; i-- {
		revCandParts[len(candidateParts)-1-i] = candidateParts[i]
	}

	idx := 0
	pat := revPats[idx]
	forceMatchEnd := true
	if strings.HasSuffix(pat, "$") {
		pat = pat[:len(pat)-1]
		forceMatchEnd = true
	}

	for i, candPart := range revCandParts {
		if hit, _ := reverseLCSImpl(pat, candPart); !hit {
			if forceMatchEnd && i == 0 {
				return false
			}
			continue
		}

		idx++
		if idx == len(revPats) {
			return true
		}
		pat = revPats[idx]
	}

	return false
}

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
