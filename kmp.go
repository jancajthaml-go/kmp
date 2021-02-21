//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
//
package kmp

// Search for pattern in given input and returns slice of matche indicies
func Search(input string, pattern string) []int {
	var (
		M = len(pattern)
		N = len(input)
	)

	if N == 0 || M == 0 {
		return nil
	}

	var (
		result = make([]int, 0)
		l = 0
		i = 1
		j = 0
		lps = make([]int, M)
	)

	for ; i < M; i++ {
		for {
			if pattern[i] == pattern[l] {
				l++
				break
			}
			if l == 0 {
				break
			}
			l = lps[l-1]
		}
		lps[i] = l
	}

	i = 0
	j = 0

	for ; i < N; i++ {
		for {
			if pattern[j] == input[i] {
				j++

				if j == M {
					result = append(result, i-j+1)
					j = lps[j-1]
				}
				break
			}
			if j > 0 {
				j = lps[j-1]
			} else {
				break
			}

		}
	}

	return result
}
