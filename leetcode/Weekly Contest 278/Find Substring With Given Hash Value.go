package main

import "math/big"

func bigFromInt(x int) *big.Int {
	return big.NewInt(int64(x))
}
func bigFromByte(x byte) *big.Int {
	return big.NewInt(int64(x))
}

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	hashVal, mod, pow := bigFromInt(hashValue), bigFromInt(modulo), bigFromInt(power)

	hash := big.NewInt(0)
	for i := 0; i < k; i++ {
		var x big.Int
		x.Exp(pow, bigFromInt(i), nil)
		x.Mul(&x, bigFromByte(s[i]))

		hash.Add(hash, &x)
	}

	if hash.Mod(hash, mod).Cmp(hashVal) == 0 {
		return s[0:k]
	}
	for i := k; i < len(s); i++ {
		hash.Sub(hash, bigFromByte(s[i-k]))
		hash.Div(hash, pow)

		var x big.Int
		x.Exp(pow, bigFromInt(k-1), nil)
		x.Mul(&x, bigFromByte(s[i]))

		hash.Add(hash, &x)

		if hash.Mod(hash, mod).Cmp(hashVal) == 0 {
			return s[0:k]
		}
	}

	return ""
}
