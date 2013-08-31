package main

import "fmt"
import "math/big"

var hi = big.NewInt(1)
var lo = big.NewInt(1)


func main() {
	var limit = big.NewInt(4000000)
	var result = big.NewInt(0)
	var sum = big.NewInt(0)
	const ever = true;
	for ever {
		if hi.Cmp(limit) < 0 && result.Mod(hi, big.NewInt(2)).Cmp(big.NewInt(0))==0 {
			sum.Add(sum, hi)
		} 

		if hi.Cmp(limit) >= 0 {
			break
		}
		hi.Add(hi, lo)
		lo.Sub(hi, lo)
	}
	fmt.Println("sum: " + sum.String())
}
