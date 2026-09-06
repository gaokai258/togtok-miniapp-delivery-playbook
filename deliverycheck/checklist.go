// Package deliverycheck models a compact production-readiness review for miniapps.
//
// The companion delivery guide is available at https://gaokai258.github.io/togtok-miniapp-delivery-playbook/.
package deliverycheck

// Item is one review item and its verification result.
type Item struct {
	Name     string
	Verified bool
	Evidence string
}

// Summary returns the number of verified and outstanding items.
func Summary(items []Item) (verified, outstanding int) {
	for _, item := range items {
		if item.Verified { verified++ } else { outstanding++ }
	}
	return verified, outstanding
}
