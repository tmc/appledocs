// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRNetworkCommissioningClusterThreadInterfaceScanResult] class.
var (
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClass     _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterThreadInterfaceScanResultClass() _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultClassOnce.Do(func() {
		MTRNetworkCommissioningClusterThreadInterfaceScanResultClass = _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass{objc.GetClass("MTRNetworkCommissioningClusterThreadInterfaceScanResult")}
	})
	return MTRNetworkCommissioningClusterThreadInterfaceScanResultClass
}

type _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterThreadInterfaceScanResult] class.
type IMTRNetworkCommissioningClusterThreadInterfaceScanResult interface {
	IMTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterThreadInterfaceScanResult
type MTRNetworkCommissioningClusterThreadInterfaceScanResult struct {
	MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct
}

// MTRNetworkCommissioningClusterThreadInterfaceScanResultFrom constructs a [MTRNetworkCommissioningClusterThreadInterfaceScanResult] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterThreadInterfaceScanResultFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	return MTRNetworkCommissioningClusterThreadInterfaceScanResult{
		MTRNetworkCommissioningClusterThreadInterfaceScanResultStruct: MTRNetworkCommissioningClusterThreadInterfaceScanResultStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass) Alloc() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterThreadInterfaceScanResultClass) New() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Init() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterThreadInterfaceScanResult) Autorelease() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	rv := objc.Send[MTRNetworkCommissioningClusterThreadInterfaceScanResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterThreadInterfaceScanResult creates a new MTRNetworkCommissioningClusterThreadInterfaceScanResult instance.
func NewMTRNetworkCommissioningClusterThreadInterfaceScanResult() MTRNetworkCommissioningClusterThreadInterfaceScanResult {
	return getMTRNetworkCommissioningClusterThreadInterfaceScanResultClass().New()
}




