// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRFixedLabelClusterLabelStruct] class.
var (
	MTRFixedLabelClusterLabelStructClass     _MTRFixedLabelClusterLabelStructClass
	MTRFixedLabelClusterLabelStructClassOnce sync.Once
)

func getMTRFixedLabelClusterLabelStructClass() _MTRFixedLabelClusterLabelStructClass {
	MTRFixedLabelClusterLabelStructClassOnce.Do(func() {
		MTRFixedLabelClusterLabelStructClass = _MTRFixedLabelClusterLabelStructClass{objc.GetClass("MTRFixedLabelClusterLabelStruct")}
	})
	return MTRFixedLabelClusterLabelStructClass
}

type _MTRFixedLabelClusterLabelStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRFixedLabelClusterLabelStruct] class.
type IMTRFixedLabelClusterLabelStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFixedLabelClusterLabelStruct
type MTRFixedLabelClusterLabelStruct struct {
	objectivec.Object
}

// MTRFixedLabelClusterLabelStructFrom constructs a [MTRFixedLabelClusterLabelStruct] from an unsafe.Pointer.
func MTRFixedLabelClusterLabelStructFrom(ptr unsafe.Pointer) MTRFixedLabelClusterLabelStruct {
	return MTRFixedLabelClusterLabelStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRFixedLabelClusterLabelStructClass) Alloc() MTRFixedLabelClusterLabelStruct {
	rv := objc.Send[MTRFixedLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRFixedLabelClusterLabelStructClass) New() MTRFixedLabelClusterLabelStruct {
	rv := objc.Send[MTRFixedLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRFixedLabelClusterLabelStruct) Init() MTRFixedLabelClusterLabelStruct {
	rv := objc.Send[MTRFixedLabelClusterLabelStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRFixedLabelClusterLabelStruct) Autorelease() MTRFixedLabelClusterLabelStruct {
	rv := objc.Send[MTRFixedLabelClusterLabelStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRFixedLabelClusterLabelStruct creates a new MTRFixedLabelClusterLabelStruct instance.
func NewMTRFixedLabelClusterLabelStruct() MTRFixedLabelClusterLabelStruct {
	return getMTRFixedLabelClusterLabelStructClass().New()
}




