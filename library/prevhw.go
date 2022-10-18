package library

// A modification of the built-in modulus % which 
// always returns a positive remainder 
func ModN(N uint, i int) int {
	m := i % int(N) 
	if m < 0 {
		m += int(N) 
	}
	return m
}

// FastPower(N,g,A) computes g^A mod N. Note that it
// assumes that N and A are non-negative. 
func FastPower(N uint, g int, A uint) int {
	var b int 
	a := g 
	b = 1 
	if A < 0 {
		A = -A 
	}
	for A > 0 {
		if A % 2 == 1 {
			b = ModN(N,b*a)
		} 
		a = ModN(N,a*a) 
		A = A / 2 
	}
	return b
}

// A separate type for the output of the Euclidean algorithm 
type EuclidData struct {
	GCD int 
	U int 
	V int
}

// Given two integers a and b, GCD(a,b) returns g,u,v where 
// g is the gcd(a,b) and au+bv = g 
func EuclidAlgo(a, b int) EuclidData {
	// if b = 0, then the gcd is a 
	if b == 0 {
		return EuclidData{a,1,0}
	}

	// Keeps tracks of the sign of a and b and makes sure 
	// a and b are non-negative 
	var na , nb   bool 
	if a < 0 {
		na = true 
		a *= -1 
	}
	if b < 0 {
		nb = true
		b *= -1 
	}

	// Variables we will return 
	var g, u, v int 
	u = 1 
	g = a
	x := 0  // keeps track of the number a's used in the Euclidean algorithm 
	y := b  // keeps track of the denominator in the Euclidean algorithm
	for y != 0 {
		t := ModN(uint(y),g) // find t,q with g = qy + t 
		q := g / y 
		s := u - q*x  
		u = x 
		g = y 
		x = s 
		y = t 
	}
	v = (g-a*u)/b  
	
	if !na && !nb {
		return EuclidData{g, u, v} 
	} else if !na && nb {
		return EuclidData{g, u, -v}
	} else if na && !nb {
		return EuclidData{g, -u, v}
	} else {
		return EuclidData{g , -u, -v}
	}
}

// Computes a such that ax = 1 mod N if gcd(a,N) = 1. Else 
// it returns -1 which serves a "junk value"

func Inverse(N uint, x int) int {
	d := EuclidAlgo(x,int(N))
	if d.GCD == 1 {
		return d.U 
	} else {
		return -1
	}
}

// Checks whether a is a Miller-Rabin witness for N being composite
func MillerRabinTest(N, a int) bool {
	// Input. Integer n to be tested, integer a as potential 
	// witness. 

	if N < 0 {
		N *= -1 
	}

	// 1. If n is even or 1 < gcd(a,n) < n, return Composite
	if N % 2 == 0 {
		return true
	} else if EuclidAlgo(N,a).GCD != 1 {
		return true
	}

	q := N-1 
	k := 0 

	// 2. Write n-1 = 2^k q with q odd
	for q % 2 != 0 {
		q = q / 2 
		k += 1 
	}

	// 3. Set a = a^q mod n. 
	a = ModN(uint(N),FastPower(uint(N),a,uint(q)))

	// 4. If a = 1 mod n, return Test Fails.
	if a == 1 {
		return false
	}

	// 5. Loop i = 0,1,2,...,k-1
	for i := 0; i < k; i++ {
	//	6. If a = -1 mod n, return Test Fails. 
		if (a + 1) % N == 0 {
			return false
		}
	//	7. Set a = a^2 mod n 
		a = FastPower(uint(N),a,2)
	}
	// 8. End i loop.

	// 9. Return Composite
	return true
}

