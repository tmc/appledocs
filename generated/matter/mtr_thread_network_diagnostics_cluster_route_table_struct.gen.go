// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClass     _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass() _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableStructClass = _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTableStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTableStruct
type MTRThreadNetworkDiagnosticsClusterRouteTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return MTRThreadNetworkDiagnosticsClusterRouteTableStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) New() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Init() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct creates a new MTRThreadNetworkDiagnosticsClusterRouteTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass().New()
}




