//go:build darwin && ios

// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for TKSmartCardTokenRegistrationManager


// Unregisters a smartcard for the provided token ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/unregisterSmartCard(tokenID:)
func (t_ TKSmartCardTokenRegistrationManager) UnregisterSmartCardWithTokenIDError(tokenID objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("unregisterSmartCardWithTokenID:error:"), tokenID, error_)
	return rv
}

// iOS-only properties

// Returns the tokenIDs of all currently registered smart card tokens
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/registeredSmartCardTokens
func (t_ TKSmartCardTokenRegistrationManager) RegisteredSmartCardTokens() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("registeredSmartCardTokens"))
	return rv
}





