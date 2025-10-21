// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BAAppExtensionInfo] class.
var (
	BAAppExtensionInfoClass     _BAAppExtensionInfoClass
	BAAppExtensionInfoClassOnce sync.Once
)

func getBAAppExtensionInfoClass() _BAAppExtensionInfoClass {
	BAAppExtensionInfoClassOnce.Do(func() {
		BAAppExtensionInfoClass = _BAAppExtensionInfoClass{objc.GetClass("BAAppExtensionInfo")}
	})
	return BAAppExtensionInfoClass
}

type _BAAppExtensionInfoClass struct {
	class objc.Class
}

// An interface definition for the [BAAppExtensionInfo] class.
type IBAAppExtensionInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAppExtensionInfo
type BAAppExtensionInfo struct {
	objectivec.Object
}

// BAAppExtensionInfoFrom constructs a [BAAppExtensionInfo] from an unsafe.Pointer.
func BAAppExtensionInfoFrom(ptr unsafe.Pointer) BAAppExtensionInfo {
	return BAAppExtensionInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BAAppExtensionInfoClass) Alloc() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BAAppExtensionInfoClass) New() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAppExtensionInfo) Init() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAppExtensionInfo) Autorelease() BAAppExtensionInfo {
	rv := objc.Send[BAAppExtensionInfo](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAppExtensionInfo creates a new BAAppExtensionInfo instance.
func NewBAAppExtensionInfo() BAAppExtensionInfo {
	return getBAAppExtensionInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAppExtensionInfo/restrictedDownloadSizeRemaining-9itic
func (b_ BAAppExtensionInfo) RestrictedDownloadSizeRemaining() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("restrictedDownloadSizeRemaining"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAppExtensionInfo/restrictedEssentialDownloadSizeRemaining-76av8
func (b_ BAAppExtensionInfo) RestrictedEssentialDownloadSizeRemaining() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("restrictedEssentialDownloadSizeRemaining"))
	return rv
}



