package dutchfees

import "testing"

func TestPegKeeperReceiptValuation(t *testing.T) {
	for _, tc := range []struct {
		name        string
		after, mask uint64
		price, want float64
		known       bool
	}{
		{"collector LP delta", 125, 0, 2, 5, true},
		{"profit paid elsewhere", 100, 0, 0, 0, true},
		{"blocked burner", 125, 4, 2, 0, true},
		{"blocked collection", 125, 2, 2, 0, true},
		{"unpriced LP", 125, 0, 0, 0, false},
		{"negative delta", 99, 0, 2, 0, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rows := []result{{true, uintWord(100)}, {true, uintWord(999)}, {true, uintWord(tc.after)}, {true, uintWord(tc.mask)}}
			got, known := pegKeeperValue(rows, tokenMeta{Decimals: 1, PriceUSD: tc.price})
			if got != tc.want || known != tc.known {
				t.Fatalf("got (%g,%t), want (%g,%t)", got, known, tc.want, tc.known)
			}
		})
	}
}
