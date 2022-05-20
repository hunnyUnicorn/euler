package LargestExponential

import (
	"math"
	"sort"
)

type ExponentValues struct {
	base     int
	exponent int
	value    float64
}

type ByValue []ExponentValues

func (pair ByValue) Len() int           { return len(pair) }
func (pair ByValue) Less(i, j int) bool { return pair[i].value > pair[j].value }
func (pair ByValue) Swap(i, j int)      { pair[i], pair[j] = pair[j], pair[i] }

func findLargestNumber(countOfPairs int, bases, exponents []int) (largestB, largestE int) {
	var values []ExponentValues
	for pairIndex := 0; pairIndex < countOfPairs; pairIndex++ {
		b := bases[pairIndex]
		e := exponents[pairIndex]
		values = append(values, ExponentValues{b, e, math.Log(float64(b)) * float64(e)})
	}
	sort.Sort(ByValue(values))
	return values[0].base, values[0].exponent
}

/*
func main() {
	var n, i, k int
	var values []ExponentValues
	fmt.Scanf("%d\n", &n)
	for i = 0; i < n; i++ {
		var b, e int
		fmt.Scanf("%d%d\n", &b, &e)
		values = append(values, ExponentValues{b, e, math.Log(float64(b)) * float64(e)})
	}
	fmt.Scanf("%d\n", &k)
	sort.Sort(ByValue(values))
	fmt.Println(values[k-1].B, values[k-1].E)
}
*/
