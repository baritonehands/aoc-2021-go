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

func SeqSet[I iter.Seq[T], T comparable](iter I) map[T]bool {
	return it.Fold(iter, func(m map[T]bool, t T) map[T]bool {
		m[t] = true
		return m
	}, make(map[T]bool))
}

func SetDifference[K comparable](lhs map[K]bool, rhs map[K]bool) map[K]bool {
	var ret = make(map[K]bool)
	for c, v := range lhs {
		_, present := rhs[c]
		if v && !present {
			ret[c] = true
		}
	}
	return ret
}

func Frequencies[I iter.Seq[T], T comparable](iter I) map[T]int {
	return it.Fold(iter, func(m map[T]int, t T) map[T]int {
		m[t]++
		return m
	}, make(map[T]int))
}

func FlatMap[V, W any, S iter.Seq[W]](delegate func(func(V) bool), f func(V) S) iter.Seq[W] {
	return func(yield func(W) bool) {
		for innerValue := range delegate {
			for value := range f(innerValue) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func FlatMap2[V, W, X, Y any, S iter.Seq2[X, Y]](delegate func(func(V, W) bool), f func(V, W) S) iter.Seq2[X, Y] {
	return func(yield func(X, Y) bool) {
		for v, w := range delegate {
			for x, y := range f(v, w) {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}
