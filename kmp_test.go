package kmp

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkKnuthMorrisPrattSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Search("123", "2")
	}
}

func BenchmarkKnuthMorrisPrattLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Search("00123014764700968325", "00")
	}
}

func BenchmarkKnuthMorrisPrattSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Search("123", "2")
		}
	})
}

func BenchmarkKnuthMorrisPrattLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Search("00123014764700968325", "00")
		}
	})
}

func TestSearch(t *testing.T) {
	expectIndicies := func(a string, b string, c []int) {
		d := Search(a, b)
		if !reflect.DeepEqual(c, d) {
			t.Errorf(fmt.Sprintf("a: `%s`, b: `%s` expected %+v, got %+v", a, b, c, d))
		}
	}

	expectIndicies("", "", nil)
	expectIndicies("", "a", nil)
	expectIndicies("a", "", nil)
	expectIndicies("lorem ipsum dolor", "lorem", []int{0})
	expectIndicies("lorem ipsum dolor ipsum", "ipsum", []int{6, 18})
	expectIndicies("lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum color lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor lorem ipsum dolor", "lorem ipsum color", []int{108})
}
