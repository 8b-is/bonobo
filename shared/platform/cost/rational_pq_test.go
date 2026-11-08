package cost

import (
	"testing"
)

func TestRationalPQ_BonoboResonance(t *testing.T) {
	bonobo := BonoboRatio()
	if bonobo.String() != "3/2" {
		t.Fatalf("expected 3/2, got %s", bonobo.String())
	}
	if bonobo.Float64() != 1.5 {
		t.Fatalf("expected 1.5, got %f", bonobo.Float64())
	}
}

func TestRationalPQ_ExactAccounting(t *testing.T) {
	// 500 tokens at $3.00 / 1M tokens
	// 500 * 3 / 1,000,000 = 1500 / 1,000,000 = 3 / 2000
	spend := ExactTokenSpendPQ(500, 3.0)
	if spend.String() != "3/2000" {
		t.Fatalf("expected 3/2000, got %s", spend.String())
	}

	// Add another 500 tokens -> 3/2000 + 3/2000 = 6/2000 = 3/1000
	spend2 := spend.Add(spend)
	if spend2.String() != "3/1000" {
		t.Fatalf("expected 3/1000, got %s", spend2.String())
	}
}
