// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] class.
var (
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass     _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass() _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClassOnce.Do(func() {
		MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass = _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass{objc.GetClass("MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct")}
	})
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass
}

type _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] class.
type IMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
type MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom constructs a [MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass) Alloc() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass) New() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Init() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct) Autorelease() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct creates a new MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct instance.
func NewMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct() MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct {
	return getMTRNetworkCommissioningClusterThreadInterfaceScanResultStructClass().New()
}




