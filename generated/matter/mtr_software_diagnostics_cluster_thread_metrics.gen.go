// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsClass     _MTRSoftwareDiagnosticsClusterThreadMetricsClass
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsClass() _MTRSoftwareDiagnosticsClusterThreadMetricsClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsClass = _MTRSoftwareDiagnosticsClusterThreadMetricsClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetrics")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
type IMTRSoftwareDiagnosticsClusterThreadMetrics interface {
	IMTRSoftwareDiagnosticsClusterThreadMetricsStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics
type MTRSoftwareDiagnosticsClusterThreadMetrics struct {
	MTRSoftwareDiagnosticsClusterThreadMetricsStruct
}

// MTRSoftwareDiagnosticsClusterThreadMetricsFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetrics] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetrics {
	return MTRSoftwareDiagnosticsClusterThreadMetrics{
		MTRSoftwareDiagnosticsClusterThreadMetricsStruct: MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) New() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Init() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetrics creates a new MTRSoftwareDiagnosticsClusterThreadMetrics instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetrics() MTRSoftwareDiagnosticsClusterThreadMetrics {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsClass().New()
}




