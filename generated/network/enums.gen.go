// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

/* debug [enums.gen.go]: Generating 1 enums for Network */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum nw_parameters_attribution_t (2 cases) */
// nw_parameters_attribution_t - The entities that can make a network request.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_attribution_t
type nw_parameters_attribution_t uint

const (
	// nw_parameters_attribution_developer - A developer-initiated network request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_attribution_t/developer
	nw_parameters_attribution_developer nw_parameters_attribution_t = 0
	// nw_parameters_attribution_user - The user explicitly directs the app to make a network request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_attribution_t/user
	nw_parameters_attribution_user nw_parameters_attribution_t = 0
)


