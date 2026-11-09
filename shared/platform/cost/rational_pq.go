package cost

import (
	"fmt"
	"math/big"
)

// RationalPQ represents an exact rational number p/q without floating-point drift.
// Built for God of Sun phylogenetic octad (Rung 0: Bonobo, p/q = 3/2).
type RationalPQ struct {
	P *big.Int `json:"p"`
	Q *big.Int `json:"q"`
}

// NewRationalPQ constructs a reduced rational p/q. If q <= 0, q defaults to 1.
func NewRationalPQ(p, q int64) RationalPQ {
	if q == 0 {
		q = 1
	}
	bp := big.NewInt(p)
	bq := big.NewInt(q)
	
	// Canonical sign on numerator
	if bq.Sign() < 0 {
		bp.Neg(bp)
		bq.Neg(bq)
	}

	gcd := new(big.Int).GCD(nil, nil, bp, bq)
	if gcd.Sign() > 0 {
		bp.Quo(bp, gcd)
		bq.Quo(bq, gcd)
	}

	return RationalPQ{P: bp, Q: bq}
}

// BonoboRatio returns the canonical Bonobo solar phylogenetic resonance ratio: 3/2 (+1 trit).
func BonoboRatio() RationalPQ {
	return NewRationalPQ(3, 2)
}

// LovaszDualRatio returns the LLL lattice basis reduction Lovasz condition parameter: 3/4.
func LovaszDualRatio() RationalPQ {
	return NewRationalPQ(3, 4)
}

// TechnoZenHarmonicRatio returns the Techno-Zen-Bonobo era Pythagorean whole-tone harmonic: (3/2) * (3/4) = 9/8.
func TechnoZenHarmonicRatio() RationalPQ {
	return BonoboRatio().Mul(LovaszDualRatio())
}

// TechnoZenSquaredRatio returns the Techno-Zen-Bonobo additive-multiplicative identity: (3/2) + (3/4) = 9/4 = (3/2)^2.
func TechnoZenSquaredRatio() RationalPQ {
	return BonoboRatio().Add(LovaszDualRatio())
}

// Div returns r / other in exact rational representation. If other == 0, returns 0/1.
func (r RationalPQ) Div(other RationalPQ) RationalPQ {
	if other.P == nil || other.P.Sign() == 0 {
		return NewRationalPQ(0, 1)
	}
	recip := RationalPQ{P: new(big.Int).Set(other.Q), Q: new(big.Int).Set(other.P)}
	if recip.Q.Sign() < 0 {
		recip.P.Neg(recip.P)
		recip.Q.Neg(recip.Q)
	}
	return r.Mul(recip)
}

// CompressionSavingsPQ returns (tokensBefore - tokensAfter) / tokensBefore in canonical reduced p/q form.
func CompressionSavingsPQ(tokensBefore, tokensAfter int64) RationalPQ {
	if tokensBefore <= 0 || tokensAfter >= tokensBefore {
		return NewRationalPQ(0, 1)
	}
	return NewRationalPQ(tokensBefore-tokensAfter, tokensBefore)
}

// Add returns r + other in exact rational representation
func (r RationalPQ) Add(other RationalPQ) RationalPQ {
	// (p1*q2 + p2*q1) / (q1*q2)
	p1q2 := new(big.Int).Mul(r.P, other.Q)
	p2q1 := new(big.Int).Mul(other.P, r.Q)
	num := new(big.Int).Add(p1q2, p2q1)
	den := new(big.Int).Mul(r.Q, other.Q)

	gcd := new(big.Int).GCD(nil, nil, num, den)
	if gcd.Sign() > 0 {
		num.Quo(num, gcd)
		den.Quo(den, gcd)
	}
	return RationalPQ{P: num, Q: den}
}

// Sub returns r - other in exact rational representation
func (r RationalPQ) Sub(other RationalPQ) RationalPQ {
	negOther := RationalPQ{P: new(big.Int).Neg(other.P), Q: other.Q}
	return r.Add(negOther)
}

// Mul returns r * other in exact rational representation
func (r RationalPQ) Mul(other RationalPQ) RationalPQ {
	num := new(big.Int).Mul(r.P, other.P)
	den := new(big.Int).Mul(r.Q, other.Q)

	gcd := new(big.Int).GCD(nil, nil, num, den)
	if gcd.Sign() > 0 {
		num.Quo(num, gcd)
		den.Quo(den, gcd)
	}
	return RationalPQ{P: num, Q: den}
}

// String formats the rational as "p/q"
func (r RationalPQ) String() string {
	if r.P == nil || r.Q == nil {
		return "0/1"
	}
	return fmt.Sprintf("%s/%s", r.P.String(), r.Q.String())
}

// Float64 returns the nearest float64 representation
func (r RationalPQ) Float64() float64 {
	if r.P == nil || r.Q == nil || r.Q.Sign() == 0 {
		return 0.0
	}
	rat := new(big.Rat).SetFrac(r.P, r.Q)
	f, _ := rat.Float64()
	return f
}

// ExactTokenSpendPQ computes exact token spend p/q given tokens and price per million (ppm/1).
// Result is tokens * ppm / 1,000,000 in canonical reduced p/q form.
func ExactTokenSpendPQ(tokens int64, pricePerMillionUSD float64) RationalPQ {
	if tokens <= 0 || pricePerMillionUSD <= 0 {
		return NewRationalPQ(0, 1)
	}
	// Convert pricePerMillionUSD to exact micro-cents or scaled integer
	// 1 USD = 1,000,000 micro-USD.
	scaledPpm := int64(pricePerMillionUSD * 100_000_000) // 8 decimal places precision
	num := big.NewInt(tokens)
	num.Mul(num, big.NewInt(scaledPpm))
	den := big.NewInt(1_000_000 * 100_000_000)

	gcd := new(big.Int).GCD(nil, nil, num, den)
	if gcd.Sign() > 0 {
		num.Quo(num, gcd)
		den.Quo(den, gcd)
	}
	return RationalPQ{P: num, Q: den}
}
