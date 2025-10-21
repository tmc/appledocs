// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] class.
var (
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass     _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass() _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass {
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass = _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents")}
	})
	return MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass
}

type _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] class.
type IMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents
type MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom constructs a [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	return MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass) Alloc() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass) New() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) Init() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) Autorelease() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents creates a new MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents instance.
func NewMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	return getMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass().New()
}




