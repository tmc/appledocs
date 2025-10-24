// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKTokenWatcherTokenInfo] class.
var (
	TKTokenWatcherTokenInfoClass     _TKTokenWatcherTokenInfoClass
	TKTokenWatcherTokenInfoClassOnce sync.Once
)

func getTKTokenWatcherTokenInfoClass() _TKTokenWatcherTokenInfoClass {
	TKTokenWatcherTokenInfoClassOnce.Do(func() {
		TKTokenWatcherTokenInfoClass = _TKTokenWatcherTokenInfoClass{objc.GetClass("TKTokenWatcherTokenInfo")}
	})
	return TKTokenWatcherTokenInfoClass
}

type _TKTokenWatcherTokenInfoClass struct {
	class objc.Class
}

// An interface definition for the [TKTokenWatcherTokenInfo] class.
type ITKTokenWatcherTokenInfo interface {
	objectivec.IObject
	// properties:
	DriverName() objc.IObject /* cross-framework: NSString */
	SlotName() objc.IObject /* cross-framework: NSString */
	TokenID() objc.IObject /* cross-framework: NSString */
	SetTokenID(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo
type TKTokenWatcherTokenInfo struct {
	objectivec.Object
}

// TKTokenWatcherTokenInfoFrom constructs a [TKTokenWatcherTokenInfo] from an unsafe.Pointer.
func TKTokenWatcherTokenInfoFrom(ptr unsafe.Pointer) TKTokenWatcherTokenInfo {
	return TKTokenWatcherTokenInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenWatcherTokenInfoClass) Alloc() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenWatcherTokenInfoClass) New() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenWatcherTokenInfo) Init() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenWatcherTokenInfo) Autorelease() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenWatcherTokenInfo creates a new TKTokenWatcherTokenInfo instance.
func NewTKTokenWatcherTokenInfo() TKTokenWatcherTokenInfo {
	return getTKTokenWatcherTokenInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/driverName
func (t_ TKTokenWatcherTokenInfo) DriverName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("driverName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/slotName
func (t_ TKTokenWatcherTokenInfo) SlotName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("slotName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenwatcher/tokeninfo/tokenid
func (t_ TKTokenWatcherTokenInfo) TokenID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("tokenID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenwatcher/tokeninfo/tokenid
func (t_ TKTokenWatcherTokenInfo) SetTokenID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenID:"), value)
}




