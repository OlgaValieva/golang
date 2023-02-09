package ex01

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

type minCoins2Test struct {
	val     int
	coins   []int
	correct []int
}

var minCoins2TestsEmpty = []minCoins2Test{
	{0, []int{1, 2, 3}, []int{}},
	{100, []int{}, []int{}},
	{0, []int{0}, []int{}},
}

func TestMinCoins2EmptyInput(t *testing.T) {
	for _, i := range minCoins2TestsEmpty {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2EmptyInput0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(minCoins2TestsEmpty[0].val, minCoins2TestsEmpty[0].coins)
	}
}

func BenchmarkMinCoins2EmptyInput1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(minCoins2TestsEmpty[1].val, minCoins2TestsEmpty[2].coins)
	}
}

func BenchmarkMinCoins2EmptyInput2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(minCoins2TestsEmpty[2].val, minCoins2TestsEmpty[2].coins)
	}
}

var MinCoins2TestsNegative = []minCoins2Test{
	{-5, []int{1, 5, 10}, []int{}},
	{10, []int{1, -5, 10}, []int{}},
	{-5, []int{1, -5, 10}, []int{}},
}

func TestMinCoins2NegativeInput(t *testing.T) {
	for _, i := range MinCoins2TestsNegative {
		timeout := time.After(3 * time.Second)
		done := make(chan []int)
		go func() {
			got := minCoins2(i.val, i.coins)
			done <- got
		}()
		select {
		case <-timeout:
			t.Fatal("Test didn't finish in time")
		case got := <-done:
			sort.Ints(got)
			if !reflect.DeepEqual(got, i.correct) {
				t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
			}
		}
	}
}

func BenchmarkMinCoins2NegativeInput0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[0].val, MinCoins2TestsNegative[0].coins)
	}
}

func BenchmarkMinCoins2NegativeInput1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[1].val, MinCoins2TestsNegative[1].coins)
	}
}

func BenchmarkMinCoins2NegativeInput2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsNegative[2].val, MinCoins2TestsNegative[2].coins)
	}
}

var MinCoins2TestsSorted = []minCoins2Test{
	{1, []int{1, 5, 10}, []int{1}},
	{5, []int{1, 5, 10}, []int{5}},
	{10, []int{1, 5, 10}, []int{10}},
}

func TestMinCoins2Sorted(t *testing.T) {
	for _, i := range MinCoins2TestsSorted {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2Sorted0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[0].val, MinCoins2TestsSorted[0].coins)
	}
}

func BenchmarkMinCoins2Sorted1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[1].val, MinCoins2TestsSorted[1].coins)
	}
}

func BenchmarkMinCoins2Sorted2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsSorted[2].val, MinCoins2TestsSorted[2].coins)
	}
}

var MinCoins2TestsUnsorted = []minCoins2Test{
	{6, []int{1, 10, 5}, []int{1, 5}},
	{6, []int{5, 1, 10}, []int{1, 5}},
	{6, []int{5, 10, 1}, []int{1, 5}},
}

func TestMinCoins2Unsorted(t *testing.T) {
	for _, i := range MinCoins2TestsUnsorted {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2Unsorted0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[0].val, MinCoins2TestsUnsorted[0].coins)
	}
}

func BenchmarkMinCoins2Unsorted1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[1].val, MinCoins2TestsUnsorted[1].coins)
	}
}

func BenchmarkMinCoins2Unsorted2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsUnsorted[2].val, MinCoins2TestsUnsorted[2].coins)
	}
}

var MinCoins2TestsImposible = []minCoins2Test{
	{4, []int{5, 10, 15}, []int{}},
	{9, []int{5, 10, 15}, []int{}},
	{14, []int{5, 10, 15}, []int{}},
}

func TestMinCoins2Imposible(t *testing.T) {
	for _, i := range MinCoins2TestsImposible {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2Imposible0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[0].val, MinCoins2TestsImposible[0].coins)
	}
}

func BenchmarkMinCoins2Imposible1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[1].val, MinCoins2TestsImposible[1].coins)
	}
}

func BenchmarkMinCoins2Imposible2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsImposible[2].val, MinCoins2TestsImposible[2].coins)
	}
}

var MinCoins2TestsDublicates = []minCoins2Test{
	{6, []int{1, 1, 5, 10}, []int{1, 5}},
	{6, []int{1, 5, 1, 10}, []int{1, 5}},
	{6, []int{1, 5, 10, 1}, []int{1, 5}},
}

func TestMinCoins2Dublicates(t *testing.T) {
	for _, i := range MinCoins2TestsDublicates {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2Dublicates0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[0].val, MinCoins2TestsDublicates[0].coins)
	}
}

func BenchmarkMinCoins2Dublicates1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[1].val, MinCoins2TestsDublicates[1].coins)
	}
}

func BenchmarkMinCoins2Dublicates2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsDublicates[2].val, MinCoins2TestsDublicates[2].coins)
	}
}

var MinCoins2TestsOptimize = []minCoins2Test{
	{15, []int{5, 1, 5, 10}, []int{5, 10}},
	{10, []int{1, 5, 10, 1}, []int{10}},
	{13, []int{5, 10, 1}, []int{1, 1, 1, 10}},
}

func TestMinCoins2Optimize(t *testing.T) {
	for _, i := range MinCoins2TestsOptimize {
		got := minCoins2(i.val, i.coins)
		sort.Ints(got)
		if !reflect.DeepEqual(got, i.correct) {
			t.Errorf("minCoins2(%d, %v) = %v, correct %v", i.val, i.coins, got, i.correct)
		}
	}
}

func BenchmarkMinCoins2Optimize0(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[0].val, MinCoins2TestsOptimize[0].coins)
	}
}

func BenchmarkMinCoins2Optimize1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[1].val, MinCoins2TestsOptimize[1].coins)
	}
}

func BenchmarkMinCoins2Optimize2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		minCoins2(MinCoins2TestsOptimize[2].val, MinCoins2TestsOptimize[2].coins)
	}
}
