// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCardTokenRegistrationManager] class.
var (
	TKSmartCardTokenRegistrationManagerClass     _TKSmartCardTokenRegistrationManagerClass
	TKSmartCardTokenRegistrationManagerClassOnce sync.Once
)

func getTKSmartCardTokenRegistrationManagerClass() _TKSmartCardTokenRegistrationManagerClass {
	TKSmartCardTokenRegistrationManagerClassOnce.Do(func() {
		TKSmartCardTokenRegistrationManagerClass = _TKSmartCardTokenRegistrationManagerClass{objc.GetClass("TKSmartCardTokenRegistrationManager")}
	})
	return TKSmartCardTokenRegistrationManagerClass
}

type _TKSmartCardTokenRegistrationManagerClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardTokenRegistrationManager] class.
type ITKSmartCardTokenRegistrationManager interface {
	objectivec.IObject
	RegisterSmartCardWithTokenIDPromptMessageError(tokenID string, promptMessage string, error_ unsafe.Pointer) bool
	UnregisterSmartCardWithTokenIDError(tokenID string, error_ unsafe.Pointer) bool
	RegisteredSmartCardTokens() []string
}

// Provides a centralized management system for registering and unregistering smartcards using their token IDs.
//
// keeps its itself accessible via Keychain and system will automatically invoke an NFC slot when a cryptographic operation is required and asks to provide the registered card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager
type TKSmartCardTokenRegistrationManager struct {
	objectivec.Object
}

// TKSmartCardTokenRegistrationManagerFrom constructs a [TKSmartCardTokenRegistrationManager] from an unsafe.Pointer.
//
// Provides a centralized management system for registering and unregistering smartcards using their token IDs.
func TKSmartCardTokenRegistrationManagerFrom(ptr unsafe.Pointer) TKSmartCardTokenRegistrationManager {
	return TKSmartCardTokenRegistrationManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenRegistrationManagerClass) Alloc() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardTokenRegistrationManagerClass) New() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenRegistrationManager) Init() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenRegistrationManager) Autorelease() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenRegistrationManager creates a new TKSmartCardTokenRegistrationManager instance.
func NewTKSmartCardTokenRegistrationManager() TKSmartCardTokenRegistrationManager {
	return getTKSmartCardTokenRegistrationManagerClass().New()
}


// Default instance of registration manager
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/default
func (tc _TKSmartCardTokenRegistrationManagerClass) DefaultManager() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("defaultManager"))
	return rv
}
// Registers a smartcard with a specific token ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/registerSmartCard(tokenID:promptMessage:)
func (t_ TKSmartCardTokenRegistrationManager) RegisterSmartCardWithTokenIDPromptMessageError(tokenID string, promptMessage string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("registerSmartCardWithTokenID:promptMessage:error:"), objc.String(tokenID), objc.String(promptMessage), error_)
	return rv
}

// Unregisters a smartcard for the provided token ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/unregisterSmartCard(tokenID:)
func (t_ TKSmartCardTokenRegistrationManager) UnregisterSmartCardWithTokenIDError(tokenID string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("unregisterSmartCardWithTokenID:error:"), objc.String(tokenID), error_)
	return rv
}

// Default instance of registration manager
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/default
func (t_ TKSmartCardTokenRegistrationManager) DefaultManager() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](t_.ID, objc.Sel("defaultManager"))
	return rv
}

// Returns the tokenIDs of all currently registered smart card tokens
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/registeredSmartCardTokens
func (t_ TKSmartCardTokenRegistrationManager) RegisteredSmartCardTokens() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("registeredSmartCardTokens"))
	return rv
}



