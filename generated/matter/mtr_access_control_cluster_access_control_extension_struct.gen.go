// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterAccessControlExtensionStruct] class.
var (
	MTRAccessControlClusterAccessControlExtensionStructClass     _MTRAccessControlClusterAccessControlExtensionStructClass
	MTRAccessControlClusterAccessControlExtensionStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlExtensionStructClass() _MTRAccessControlClusterAccessControlExtensionStructClass {
	MTRAccessControlClusterAccessControlExtensionStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlExtensionStructClass = _MTRAccessControlClusterAccessControlExtensionStructClass{objc.GetClass("MTRAccessControlClusterAccessControlExtensionStruct")}
	})
	return MTRAccessControlClusterAccessControlExtensionStructClass
}

type _MTRAccessControlClusterAccessControlExtensionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlExtensionStruct] class.
type IMTRAccessControlClusterAccessControlExtensionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlExtensionStruct
type MTRAccessControlClusterAccessControlExtensionStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlExtensionStructFrom constructs a [MTRAccessControlClusterAccessControlExtensionStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlExtensionStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlExtensionStruct {
	return MTRAccessControlClusterAccessControlExtensionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlExtensionStructClass) Alloc() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlExtensionStructClass) New() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) Init() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) Autorelease() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlExtensionStruct creates a new MTRAccessControlClusterAccessControlExtensionStruct instance.
func NewMTRAccessControlClusterAccessControlExtensionStruct() MTRAccessControlClusterAccessControlExtensionStruct {
	return getMTRAccessControlClusterAccessControlExtensionStructClass().New()
}




