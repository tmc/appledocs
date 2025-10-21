// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRFabricInfo] class.
var (
	MTRFabricInfoClass     _MTRFabricInfoClass
	MTRFabricInfoClassOnce sync.Once
)

func getMTRFabricInfoClass() _MTRFabricInfoClass {
	MTRFabricInfoClassOnce.Do(func() {
		MTRFabricInfoClass = _MTRFabricInfoClass{objc.GetClass("MTRFabricInfo")}
	})
	return MTRFabricInfoClass
}

type _MTRFabricInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRFabricInfo] class.
type IMTRFabricInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo
type MTRFabricInfo struct {
	objectivec.Object
}

// MTRFabricInfoFrom constructs a [MTRFabricInfo] from an unsafe.Pointer.
func MTRFabricInfoFrom(ptr unsafe.Pointer) MTRFabricInfo {
	return MTRFabricInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRFabricInfoClass) Alloc() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRFabricInfoClass) New() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRFabricInfo) Init() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRFabricInfo) Autorelease() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRFabricInfo creates a new MTRFabricInfo instance.
func NewMTRFabricInfo() MTRFabricInfo {
	return getMTRFabricInfoClass().New()
}




