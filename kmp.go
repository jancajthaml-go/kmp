//
// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
//
package kmp

func Search(input, pattern string) []int {
	var M = len(pattern)
	var N = len(input)

	if N == 0 || M == 0 {
		return nil
	}

	var result = make([]int, 0)
	var l = 0
	var i = 0
	var j = 0

	var lps = make([]int, M)

	for i = 1; i < M; i++ {
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

	for i, j = 0, 0; i < N; i++ {
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
