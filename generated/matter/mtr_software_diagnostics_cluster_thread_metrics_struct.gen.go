// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClass     _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass() _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsStructClass = _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetricsStruct")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
type IMTRSoftwareDiagnosticsClusterThreadMetricsStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct
type MTRSoftwareDiagnosticsClusterThreadMetricsStruct struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return MTRSoftwareDiagnosticsClusterThreadMetricsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) New() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Init() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct creates a new MTRSoftwareDiagnosticsClusterThreadMetricsStruct instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass().New()
}




