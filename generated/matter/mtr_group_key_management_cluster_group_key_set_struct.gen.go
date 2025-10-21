// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGroupKeyManagementClusterGroupKeySetStruct] class.
var (
	MTRGroupKeyManagementClusterGroupKeySetStructClass     _MTRGroupKeyManagementClusterGroupKeySetStructClass
	MTRGroupKeyManagementClusterGroupKeySetStructClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterGroupKeySetStructClass() _MTRGroupKeyManagementClusterGroupKeySetStructClass {
	MTRGroupKeyManagementClusterGroupKeySetStructClassOnce.Do(func() {
		MTRGroupKeyManagementClusterGroupKeySetStructClass = _MTRGroupKeyManagementClusterGroupKeySetStructClass{objc.GetClass("MTRGroupKeyManagementClusterGroupKeySetStruct")}
	})
	return MTRGroupKeyManagementClusterGroupKeySetStructClass
}

type _MTRGroupKeyManagementClusterGroupKeySetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterGroupKeySetStruct] class.
type IMTRGroupKeyManagementClusterGroupKeySetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterGroupKeySetStruct
type MTRGroupKeyManagementClusterGroupKeySetStruct struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterGroupKeySetStructFrom constructs a [MTRGroupKeyManagementClusterGroupKeySetStruct] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterGroupKeySetStructFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterGroupKeySetStruct {
	return MTRGroupKeyManagementClusterGroupKeySetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterGroupKeySetStructClass) Alloc() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterGroupKeySetStructClass) New() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) Init() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) Autorelease() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterGroupKeySetStruct creates a new MTRGroupKeyManagementClusterGroupKeySetStruct instance.
func NewMTRGroupKeyManagementClusterGroupKeySetStruct() MTRGroupKeyManagementClusterGroupKeySetStruct {
	return getMTRGroupKeyManagementClusterGroupKeySetStructClass().New()
}




