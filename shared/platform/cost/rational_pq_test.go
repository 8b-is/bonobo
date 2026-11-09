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

	lovasz := LovaszDualRatio()
	if lovasz.String() != "3/4" {
		t.Fatalf("expected 3/4, got %s", lovasz.String())
	}

	harmonic := TechnoZenHarmonicRatio()
	if harmonic.String() != "9/8" {
		t.Fatalf("expected 9/8, got %s", harmonic.String())
	}

	squared := TechnoZenSquaredRatio()
	if squared.String() != "9/4" {
		t.Fatalf("expected 9/4, got %s", squared.String())
	}

	octave := bonobo.Div(lovasz)
	if octave.String() != "2/1" {
		t.Fatalf("expected 2/1, got %s", octave.String())
	}

	savings := CompressionSavingsPQ(80, 28)
	if savings.String() != "13/20" {
		t.Fatalf("expected 13/20, got %s", savings.String())
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
