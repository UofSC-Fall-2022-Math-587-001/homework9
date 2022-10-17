package hw8

import (
	"errors"
	// "github.com/UofSC-Fall-2022-Math-587-001/homework8/library"
)

func Pollard(N, b int) (int,error) {
	// Implement Pollard's p-1 factorization test. 
	// Input: N an integer to be factored and b a bound for 
	// the exponents 
	// 1. Set a = 2 (or some other convenient value)
	// 2. Loop j = 2,3,4,...,b 
	// 	3. Set a = a^j mod N 
	//	4. Compute d = gcd(a-1,N) 
	//	5. If 1 < d < N, then return d and nil for the error 
	// 6. Return 0 and "Test failed" for the error
	return 0, errors.New("Test Failed") 
}

