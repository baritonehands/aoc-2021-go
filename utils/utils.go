package utils

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"iter"
	"strings"
)

func Split2(s string) (string, string) {
	arr := strings.SplitN(s, " ", 2)
	return arr[0], arr[1]
}

func Frequencies[I iter.Seq[T], T comparable](iter I) map[T]int {
	return it.Fold(iter, func(m map[T]int, t T) map[T]int {
		m[t]++
		return m
	}, make(map[T]int))
}

func Concat2[T any](lhs iter.Seq[T], rhs iter.Seq[T]) iter.Seq[T] {
	return it.Chain(lhs, rhs)
}
