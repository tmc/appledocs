// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] class.
var (
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass     _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterWiFiInterfaceScanResultClass() _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultClassOnce.Do(func() {
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass = _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass{objc.GetClass("MTRNetworkCommissioningClusterWiFiInterfaceScanResult")}
	})
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass
}

type _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] class.
type IMTRNetworkCommissioningClusterWiFiInterfaceScanResult interface {
	IMTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterWiFiInterfaceScanResult
type MTRNetworkCommissioningClusterWiFiInterfaceScanResult struct {
	MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct
}

// MTRNetworkCommissioningClusterWiFiInterfaceScanResultFrom constructs a [MTRNetworkCommissioningClusterWiFiInterfaceScanResult] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterWiFiInterfaceScanResultFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	return MTRNetworkCommissioningClusterWiFiInterfaceScanResult{
		MTRNetworkCommissioningClusterWiFiInterfaceScanResultStruct: MTRNetworkCommissioningClusterWiFiInterfaceScanResultStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass) Alloc() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterWiFiInterfaceScanResultClass) New() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Init() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterWiFiInterfaceScanResult) Autorelease() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterWiFiInterfaceScanResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterWiFiInterfaceScanResult creates a new MTRNetworkCommissioningClusterWiFiInterfaceScanResult instance.
func NewMTRNetworkCommissioningClusterWiFiInterfaceScanResult() MTRNetworkCommissioningClusterWiFiInterfaceScanResult {
	return getMTRNetworkCommissioningClusterWiFiInterfaceScanResultClass().New()
}




