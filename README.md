## Knuth-Morris-Pratt algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/kmp)](https://goreportcard.com/report/jancajthaml-go/kmp)

### Usage ###

```
import "github.com/jancajthaml-go/kmp"

kmp.Search("GEEKS FOR GEEKS", "GEEK")
```

### Performance ###

```
BenchmarkHammingParallel-4    300000000           4.25 ns/op
BenchmarkHammingSerial-4      200000000           8.77 ns/op
```

test on your own by running `make benchmark`

## Resources

* [Wikipedia - Knuth-Morris-Pratt algorithm](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)
