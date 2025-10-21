// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/driverName
func (t_ TKTokenWatcherTokenInfo) DriverName() string {
	rv := objc.Send[string](t_.ID, objc.Sel("driverName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/slotName
func (t_ TKTokenWatcherTokenInfo) SlotName() string {
	rv := objc.Send[string](t_.ID, objc.Sel("slotName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/tokenID
func (t_ TKTokenWatcherTokenInfo) TokenID() string {
	rv := objc.Send[string](t_.ID, objc.Sel("tokenID"))
	return rv
}



