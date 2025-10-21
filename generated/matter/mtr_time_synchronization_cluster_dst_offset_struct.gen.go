// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterDSTOffsetStruct] class.
var (
	MTRTimeSynchronizationClusterDSTOffsetStructClass     _MTRTimeSynchronizationClusterDSTOffsetStructClass
	MTRTimeSynchronizationClusterDSTOffsetStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTOffsetStructClass() _MTRTimeSynchronizationClusterDSTOffsetStructClass {
	MTRTimeSynchronizationClusterDSTOffsetStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTOffsetStructClass = _MTRTimeSynchronizationClusterDSTOffsetStructClass{objc.GetClass("MTRTimeSynchronizationClusterDSTOffsetStruct")}
	})
	return MTRTimeSynchronizationClusterDSTOffsetStructClass
}

type _MTRTimeSynchronizationClusterDSTOffsetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterDSTOffsetStruct] class.
type IMTRTimeSynchronizationClusterDSTOffsetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTOffsetStruct
type MTRTimeSynchronizationClusterDSTOffsetStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTOffsetStructFrom constructs a [MTRTimeSynchronizationClusterDSTOffsetStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTOffsetStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTOffsetStruct {
	return MTRTimeSynchronizationClusterDSTOffsetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTOffsetStructClass) Alloc() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterDSTOffsetStructClass) New() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) Init() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTOffsetStruct) Autorelease() MTRTimeSynchronizationClusterDSTOffsetStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTOffsetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTOffsetStruct creates a new MTRTimeSynchronizationClusterDSTOffsetStruct instance.
func NewMTRTimeSynchronizationClusterDSTOffsetStruct() MTRTimeSynchronizationClusterDSTOffsetStruct {
	return getMTRTimeSynchronizationClusterDSTOffsetStructClass().New()
}




