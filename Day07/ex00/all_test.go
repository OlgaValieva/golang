package ex00

import (
	"reflect"
	"testing"
)

type addTest struct {
	val     int
	coins   []int
	correct []int
}

var TestVars = []addTest{
	addTest{val: 50, coins: []int{1, 5, 10, 50, 100, 500, 1000}, correct: []int{50}},
	addTest{val: 5, coins: []int{1, 3, 4, 7, 13, 15}, correct: []int{4, 1}},
	addTest{val: 13, coins: []int{1, 5, 10}, correct: []int{10, 1, 1, 1}},
	addTest{val: 5, coins: []int{1, 1, 1, 1, 1, 1}, correct: []int{1, 1, 1, 1, 1}},
	addTest{val: 6, coins: []int{}, correct: []int{}},

	addTest{val: 6, coins: []int{1, 2, 6, 9, 5}, correct: []int{6}},
	addTest{val: 6, coins: []int{1, 2, 6, 1, 1}, correct: []int{6}},
	addTest{val: 2, coins: []int{1, 2, 6, 1, 1}, correct: []int{2}},
	addTest{val: 25, coins: []int{5, 10, 24}, correct: []int{10, 10, 5}},
	addTest{val: 13, coins: []int{13, 5, 10}, correct: []int{13}},
}

func TestMinCoins(t *testing.T) {
	for _, test := range TestVars {
		got := minCoins(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("minCoins(%d, %v) got = %v it should give you %v", test.val, test.coins, got, want)
		}
	}
}

func TestMinCoins2(t *testing.T) {
	for _, test := range TestVars {
		got := minCoins2(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("minCoins2(%d, %v) got = %v it should give you %v", test.val, test.coins, got, want)
		}
	}
}
