// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBasicClusterCapabilityMinimaStruct] class.
var (
	MTRBasicClusterCapabilityMinimaStructClass     _MTRBasicClusterCapabilityMinimaStructClass
	MTRBasicClusterCapabilityMinimaStructClassOnce sync.Once
)

func getMTRBasicClusterCapabilityMinimaStructClass() _MTRBasicClusterCapabilityMinimaStructClass {
	MTRBasicClusterCapabilityMinimaStructClassOnce.Do(func() {
		MTRBasicClusterCapabilityMinimaStructClass = _MTRBasicClusterCapabilityMinimaStructClass{objc.GetClass("MTRBasicClusterCapabilityMinimaStruct")}
	})
	return MTRBasicClusterCapabilityMinimaStructClass
}

type _MTRBasicClusterCapabilityMinimaStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterCapabilityMinimaStruct] class.
type IMTRBasicClusterCapabilityMinimaStruct interface {
	IMTRBasicInformationClusterCapabilityMinimaStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterCapabilityMinimaStruct
type MTRBasicClusterCapabilityMinimaStruct struct {
	MTRBasicInformationClusterCapabilityMinimaStruct
}

// MTRBasicClusterCapabilityMinimaStructFrom constructs a [MTRBasicClusterCapabilityMinimaStruct] from an unsafe.Pointer.
func MTRBasicClusterCapabilityMinimaStructFrom(ptr unsafe.Pointer) MTRBasicClusterCapabilityMinimaStruct {
	return MTRBasicClusterCapabilityMinimaStruct{
		MTRBasicInformationClusterCapabilityMinimaStruct: MTRBasicInformationClusterCapabilityMinimaStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterCapabilityMinimaStructClass) Alloc() MTRBasicClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicClusterCapabilityMinimaStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterCapabilityMinimaStructClass) New() MTRBasicClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicClusterCapabilityMinimaStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterCapabilityMinimaStruct) Init() MTRBasicClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicClusterCapabilityMinimaStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterCapabilityMinimaStruct) Autorelease() MTRBasicClusterCapabilityMinimaStruct {
	rv := objc.Send[MTRBasicClusterCapabilityMinimaStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterCapabilityMinimaStruct creates a new MTRBasicClusterCapabilityMinimaStruct instance.
func NewMTRBasicClusterCapabilityMinimaStruct() MTRBasicClusterCapabilityMinimaStruct {
	return getMTRBasicClusterCapabilityMinimaStructClass().New()
}




