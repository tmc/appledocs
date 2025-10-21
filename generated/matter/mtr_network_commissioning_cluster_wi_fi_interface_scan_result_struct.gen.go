// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] class.
var (
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass     _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass() _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClassOnce.Do(func() {
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass = _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass{objc.GetClass("MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct")}
	})
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass
}

type _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] class.
type IMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
type MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom constructs a [MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass) Alloc() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass) New() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Init() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct) Autorelease() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct creates a new MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct instance.
func NewMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct() MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct {
	return getMTRNetworkCommissioningClusterWiFiInterfaceScanResultStructClass().New()
}




