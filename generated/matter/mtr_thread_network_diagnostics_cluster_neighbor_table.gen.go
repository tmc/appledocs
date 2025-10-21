// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTable")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTable interface {
	IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable
type MTRThreadNetworkDiagnosticsClusterNeighborTable struct {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStruct
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTable] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNeighborTableFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return MTRThreadNetworkDiagnosticsClusterNeighborTable{
		MTRThreadNetworkDiagnosticsClusterNeighborTableStruct: MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Init() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTable creates a new MTRThreadNetworkDiagnosticsClusterNeighborTable instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTable() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableClass().New()
}




