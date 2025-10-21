// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTable] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableClass     _MTRThreadNetworkDiagnosticsClusterRouteTableClass
	MTRThreadNetworkDiagnosticsClusterRouteTableClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableClass() _MTRThreadNetworkDiagnosticsClusterRouteTableClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableClass = _MTRThreadNetworkDiagnosticsClusterRouteTableClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTable")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTable] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTable interface {
	IMTRThreadNetworkDiagnosticsClusterRouteTableStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable
type MTRThreadNetworkDiagnosticsClusterRouteTable struct {
	MTRThreadNetworkDiagnosticsClusterRouteTableStruct
}

// MTRThreadNetworkDiagnosticsClusterRouteTableFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTable] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterRouteTableFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTable {
	return MTRThreadNetworkDiagnosticsClusterRouteTable{
		MTRThreadNetworkDiagnosticsClusterRouteTableStruct: MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableClass) New() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Init() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTable creates a new MTRThreadNetworkDiagnosticsClusterRouteTable instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTable() MTRThreadNetworkDiagnosticsClusterRouteTable {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableClass().New()
}




