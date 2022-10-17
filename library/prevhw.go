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

// Given two integers a and b, GCD(a,b) returns g,u,v where 
// g is the gcd(a,b) and au+bv = g 
func GCD(a, b int) (int,int,int) {
	// if b = 0, then the gcd is a 
	if b == 0 {
		return a,1,0
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
		return g, u, v 
	} else if !na && nb {
		return g, u, -v 
	} else if na && !nb {
		return g, -u, v 
	} else {
		return g , -u, -v 
	}
}

// Computes a such that ax = 1 mod N if gcd(a,N) = 1. Else 
// it returns -1 which serves a "junk value"
func Inverse(N uint, x int) int {
	gcd, a, _ := GCD(x,int(N))
	if gcd == 1 {
		return a 
	} else {
		return -1
	}
}

