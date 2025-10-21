// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRServerAttribute] class.
var (
	MTRServerAttributeClass     _MTRServerAttributeClass
	MTRServerAttributeClassOnce sync.Once
)

func getMTRServerAttributeClass() _MTRServerAttributeClass {
	MTRServerAttributeClassOnce.Do(func() {
		MTRServerAttributeClass = _MTRServerAttributeClass{objc.GetClass("MTRServerAttribute")}
	})
	return MTRServerAttributeClass
}

type _MTRServerAttributeClass struct {
	class objc.Class
}

// An interface definition for the [MTRServerAttribute] class.
type IMTRServerAttribute interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServerAttribute
type MTRServerAttribute struct {
	objectivec.Object
}

// MTRServerAttributeFrom constructs a [MTRServerAttribute] from an unsafe.Pointer.
func MTRServerAttributeFrom(ptr unsafe.Pointer) MTRServerAttribute {
	return MTRServerAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServerAttributeClass) Alloc() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServerAttributeClass) New() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServerAttribute) Init() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServerAttribute) Autorelease() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServerAttribute creates a new MTRServerAttribute instance.
func NewMTRServerAttribute() MTRServerAttribute {
	return getMTRServerAttributeClass().New()
}




