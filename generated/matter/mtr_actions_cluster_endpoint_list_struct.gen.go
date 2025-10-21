// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRActionsClusterEndpointListStruct] class.
var (
	MTRActionsClusterEndpointListStructClass     _MTRActionsClusterEndpointListStructClass
	MTRActionsClusterEndpointListStructClassOnce sync.Once
)

func getMTRActionsClusterEndpointListStructClass() _MTRActionsClusterEndpointListStructClass {
	MTRActionsClusterEndpointListStructClassOnce.Do(func() {
		MTRActionsClusterEndpointListStructClass = _MTRActionsClusterEndpointListStructClass{objc.GetClass("MTRActionsClusterEndpointListStruct")}
	})
	return MTRActionsClusterEndpointListStructClass
}

type _MTRActionsClusterEndpointListStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterEndpointListStruct] class.
type IMTRActionsClusterEndpointListStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct
type MTRActionsClusterEndpointListStruct struct {
	objectivec.Object
}

// MTRActionsClusterEndpointListStructFrom constructs a [MTRActionsClusterEndpointListStruct] from an unsafe.Pointer.
func MTRActionsClusterEndpointListStructFrom(ptr unsafe.Pointer) MTRActionsClusterEndpointListStruct {
	return MTRActionsClusterEndpointListStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEndpointListStructClass) Alloc() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterEndpointListStructClass) New() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEndpointListStruct) Init() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEndpointListStruct) Autorelease() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEndpointListStruct creates a new MTRActionsClusterEndpointListStruct instance.
func NewMTRActionsClusterEndpointListStruct() MTRActionsClusterEndpointListStruct {
	return getMTRActionsClusterEndpointListStructClass().New()
}




