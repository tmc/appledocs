// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDistinguishedNameInfo] class.
var (
	MTRDistinguishedNameInfoClass     _MTRDistinguishedNameInfoClass
	MTRDistinguishedNameInfoClassOnce sync.Once
)

func getMTRDistinguishedNameInfoClass() _MTRDistinguishedNameInfoClass {
	MTRDistinguishedNameInfoClassOnce.Do(func() {
		MTRDistinguishedNameInfoClass = _MTRDistinguishedNameInfoClass{objc.GetClass("MTRDistinguishedNameInfo")}
	})
	return MTRDistinguishedNameInfoClass
}

type _MTRDistinguishedNameInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRDistinguishedNameInfo] class.
type IMTRDistinguishedNameInfo interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo
type MTRDistinguishedNameInfo struct {
	objectivec.Object
}

// MTRDistinguishedNameInfoFrom constructs a [MTRDistinguishedNameInfo] from an unsafe.Pointer.
func MTRDistinguishedNameInfoFrom(ptr unsafe.Pointer) MTRDistinguishedNameInfo {
	return MTRDistinguishedNameInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDistinguishedNameInfoClass) Alloc() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDistinguishedNameInfoClass) New() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDistinguishedNameInfo) Init() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDistinguishedNameInfo) Autorelease() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDistinguishedNameInfo creates a new MTRDistinguishedNameInfo instance.
func NewMTRDistinguishedNameInfo() MTRDistinguishedNameInfo {
	return getMTRDistinguishedNameInfoClass().New()
}




