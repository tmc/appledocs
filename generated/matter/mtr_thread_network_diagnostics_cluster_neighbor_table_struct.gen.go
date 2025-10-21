// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTableStruct
type MTRThreadNetworkDiagnosticsClusterNeighborTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Init() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct creates a new MTRThreadNetworkDiagnosticsClusterNeighborTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass().New()
}




