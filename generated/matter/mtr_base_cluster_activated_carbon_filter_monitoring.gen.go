// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterActivatedCarbonFilterMonitoring] class.
var (
	MTRBaseClusterActivatedCarbonFilterMonitoringClass     _MTRBaseClusterActivatedCarbonFilterMonitoringClass
	MTRBaseClusterActivatedCarbonFilterMonitoringClassOnce sync.Once
)

func getMTRBaseClusterActivatedCarbonFilterMonitoringClass() _MTRBaseClusterActivatedCarbonFilterMonitoringClass {
	MTRBaseClusterActivatedCarbonFilterMonitoringClassOnce.Do(func() {
		MTRBaseClusterActivatedCarbonFilterMonitoringClass = _MTRBaseClusterActivatedCarbonFilterMonitoringClass{objc.GetClass("MTRBaseClusterActivatedCarbonFilterMonitoring")}
	})
	return MTRBaseClusterActivatedCarbonFilterMonitoringClass
}

type _MTRBaseClusterActivatedCarbonFilterMonitoringClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterActivatedCarbonFilterMonitoring] class.
type IMTRBaseClusterActivatedCarbonFilterMonitoring interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterActivatedCarbonFilterMonitoring
type MTRBaseClusterActivatedCarbonFilterMonitoring struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterActivatedCarbonFilterMonitoringFrom constructs a [MTRBaseClusterActivatedCarbonFilterMonitoring] from an unsafe.Pointer.
func MTRBaseClusterActivatedCarbonFilterMonitoringFrom(ptr unsafe.Pointer) MTRBaseClusterActivatedCarbonFilterMonitoring {
	return MTRBaseClusterActivatedCarbonFilterMonitoring{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterActivatedCarbonFilterMonitoringClass) Alloc() MTRBaseClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRBaseClusterActivatedCarbonFilterMonitoring](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterActivatedCarbonFilterMonitoringClass) New() MTRBaseClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRBaseClusterActivatedCarbonFilterMonitoring](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterActivatedCarbonFilterMonitoring) Init() MTRBaseClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRBaseClusterActivatedCarbonFilterMonitoring](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterActivatedCarbonFilterMonitoring) Autorelease() MTRBaseClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRBaseClusterActivatedCarbonFilterMonitoring](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterActivatedCarbonFilterMonitoring creates a new MTRBaseClusterActivatedCarbonFilterMonitoring instance.
func NewMTRBaseClusterActivatedCarbonFilterMonitoring() MTRBaseClusterActivatedCarbonFilterMonitoring {
	return getMTRBaseClusterActivatedCarbonFilterMonitoringClass().New()
}




