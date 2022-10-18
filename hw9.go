package hw9

import (
	"math"
	"github.com/UofSC-Fall-2022-Math-587-001/homework9/library"
)

// Computes the primes <= B. Is correct assuming the validity of the 
// extended Riemann hypothesis (ERH). If you don't believe the ERH, then 
// you can change the upper bound to n/4+1, which boosts complexity 
// to exponential in the number of bits of n
func FactorBase(B int) []int {
	var primes []int 
	for n := 2; n <= B; n++ {
		// a consequence of the ERH is that we have the following
		// upper bound on a Miller-Rabin witness for n 
		upperbound := 2*math.Log(float64(n))*math.Log(float64(n))
		// for small numbers, this upperbound is a little too big
		upperbound = math.Min(float64(n),upperbound-1)
		composite := false // keeps track of non-primalty  
		for a := 2; a <= int(upperbound); a++ {
			// evaluate the Miller-Rabin test with n and a 
			b := library.MillerRabinTest(n,a) 
			// if it returns true (= composite) then mark 
			// n as composite and quit the for loop
			if b {
				composite = true 
				break
			}
		}
		// if we made it through all the possible MR witnesses 
		// without turning composite then we are prime 
		if !composite {
			primes = append(primes, n)
		}
	}
	return primes
}

func IsBSmooth(n int,l []int) bool {
	// Return true if n is B-smooth and false if not 
	return false
}

func Psi(B, N int) []int {
	// Computes the number of B-smooth integers <= N 
	var bsmooth []int 
	primes := FactorBase(B)
	for i := 2; i <= N; i++ {
		// i is B-smooth for our factor base then add it 
		// to the slice
		if IsBSmooth(i,primes) {
			bsmooth = append(bsmooth, i)
		}
	}	
	return bsmooth
}

