package deliverycheck

import "testing"

func TestSummary(t *testing.T) {
	verified, outstanding := Summary([]Item{{Verified: true}, {Verified: false}})
	if verified != 1 || outstanding != 1 { t.Fatalf("got %d verified and %d outstanding", verified, outstanding) }
}
