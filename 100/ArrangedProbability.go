package ArrangedProbability

import (
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func mod(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mod(a, b)
}

func solve(p, q, d int) (*big.Int, *big.Int) {
	P := big.NewInt(int64(p))
	Q := big.NewInt(int64(q))
	D := big.NewInt(int64(d))
	limit := big.NewInt(1)
	for i := 1; i <= 63; i++ {
		limit = mul(limit, big.NewInt(2))
	}
	for totalCountofDisk := add(big.NewInt(0), D); true && totalCountofDisk.Cmp(limit) < 0; totalCountofDisk = add(totalCountofDisk, big.NewInt(1)) {
		// b*(b-1) = n*(n-1) * q/p
		rightSide := mul(mul(totalCountofDisk, sub(totalCountofDisk, big.NewInt(1))), P)
		if mod(rightSide, Q).Cmp(big.NewInt(0)) != 0 {
			continue
		}
		rightSide = div(rightSide, Q)
		// b*(b-1) = s
		// b = roof of sqrt(s)
		countofBlueDisk := add(big.NewInt(0).Sqrt(rightSide), big.NewInt(1))

		if mul(countofBlueDisk, sub(countofBlueDisk, big.NewInt(1))).Cmp(rightSide) != 0 {
			continue
		}

		return countofBlueDisk, totalCountofDisk
	}
	return big.NewInt(0), big.NewInt(0)
}

// func main() {
// 	var t, i int
// 	pArray := make(map[int]int)
// 	qArray := make(map[int]int)
// 	dArray := make(map[int]int)
// 	fmt.Scanf("%d\n", &t)
// 	for i = 0; i < t; i++ {
// 		var p, q, d int
// 		fmt.Scanf("%d%d%d\n", &p, &q, &d)
// 		pArray[i] = p
// 		qArray[i] = q
// 		dArray[i] = d
// 	}
// 	for i = 0; i < t; i++ {
// 		solve(pArray[i], qArray[i], dArray[i])
// 	}
// }
