// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROptionalQRCodeInfo] class.
var (
	MTROptionalQRCodeInfoClass     _MTROptionalQRCodeInfoClass
	MTROptionalQRCodeInfoClassOnce sync.Once
)

func getMTROptionalQRCodeInfoClass() _MTROptionalQRCodeInfoClass {
	MTROptionalQRCodeInfoClassOnce.Do(func() {
		MTROptionalQRCodeInfoClass = _MTROptionalQRCodeInfoClass{objc.GetClass("MTROptionalQRCodeInfo")}
	})
	return MTROptionalQRCodeInfoClass
}

type _MTROptionalQRCodeInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTROptionalQRCodeInfo] class.
type IMTROptionalQRCodeInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROptionalQRCodeInfo
type MTROptionalQRCodeInfo struct {
	objectivec.Object
}

// MTROptionalQRCodeInfoFrom constructs a [MTROptionalQRCodeInfo] from an unsafe.Pointer.
func MTROptionalQRCodeInfoFrom(ptr unsafe.Pointer) MTROptionalQRCodeInfo {
	return MTROptionalQRCodeInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROptionalQRCodeInfoClass) Alloc() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROptionalQRCodeInfoClass) New() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROptionalQRCodeInfo) Init() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROptionalQRCodeInfo) Autorelease() MTROptionalQRCodeInfo {
	rv := objc.Send[MTROptionalQRCodeInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROptionalQRCodeInfo creates a new MTROptionalQRCodeInfo instance.
func NewMTROptionalQRCodeInfo() MTROptionalQRCodeInfo {
	return getMTROptionalQRCodeInfoClass().New()
}




