## Knuth-Morris-Pratt algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/kmp)](https://goreportcard.com/report/jancajthaml-go/kmp)

### Usage ###

```
import "github.com/jancajthaml-go/kmp"

kmp.Search("GEEKS FOR GEEKS", "GEEK")
```

### Performance ###

```
BenchmarkKnuthMorrisPrattSmall-4          20000000  63.6 ns/op  16 B/op  2 allocs/op
BenchmarkKnuthMorrisPrattLarge-4          10000000   153 ns/op  40 B/op  3 allocs/op
BenchmarkKnuthMorrisPrattSmallParallel-4  50000000  37.8 ns/op  16 B/op  2 allocs/op
BenchmarkKnuthMorrisPrattLargeParallel-4  20000000  86.1 ns/op  40 B/op  3 allocs/op
```

verify your performance by running `make benchmark`

## Resources

* [Wikipedia - Knuth-Morris-Pratt algorithm](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)
