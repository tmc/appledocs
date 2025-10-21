// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterGroupKeyMapStruct] class.
var (
	MTRGroupKeyManagementClusterGroupKeyMapStructClass     _MTRGroupKeyManagementClusterGroupKeyMapStructClass
	MTRGroupKeyManagementClusterGroupKeyMapStructClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterGroupKeyMapStructClass() _MTRGroupKeyManagementClusterGroupKeyMapStructClass {
	MTRGroupKeyManagementClusterGroupKeyMapStructClassOnce.Do(func() {
		MTRGroupKeyManagementClusterGroupKeyMapStructClass = _MTRGroupKeyManagementClusterGroupKeyMapStructClass{objc.GetClass("MTRGroupKeyManagementClusterGroupKeyMapStruct")}
	})
	return MTRGroupKeyManagementClusterGroupKeyMapStructClass
}

type _MTRGroupKeyManagementClusterGroupKeyMapStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterGroupKeyMapStruct] class.
type IMTRGroupKeyManagementClusterGroupKeyMapStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterGroupKeyMapStruct
type MTRGroupKeyManagementClusterGroupKeyMapStruct struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterGroupKeyMapStructFrom constructs a [MTRGroupKeyManagementClusterGroupKeyMapStruct] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterGroupKeyMapStructFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterGroupKeyMapStruct {
	return MTRGroupKeyManagementClusterGroupKeyMapStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterGroupKeyMapStructClass) Alloc() MTRGroupKeyManagementClusterGroupKeyMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeyMapStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterGroupKeyMapStructClass) New() MTRGroupKeyManagementClusterGroupKeyMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeyMapStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterGroupKeyMapStruct) Init() MTRGroupKeyManagementClusterGroupKeyMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeyMapStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterGroupKeyMapStruct) Autorelease() MTRGroupKeyManagementClusterGroupKeyMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeyMapStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterGroupKeyMapStruct creates a new MTRGroupKeyManagementClusterGroupKeyMapStruct instance.
func NewMTRGroupKeyManagementClusterGroupKeyMapStruct() MTRGroupKeyManagementClusterGroupKeyMapStruct {
	return getMTRGroupKeyManagementClusterGroupKeyMapStructClass().New()
}




