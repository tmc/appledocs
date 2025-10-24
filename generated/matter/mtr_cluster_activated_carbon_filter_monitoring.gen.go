// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterActivatedCarbonFilterMonitoring] class.
var (
	MTRClusterActivatedCarbonFilterMonitoringClass     _MTRClusterActivatedCarbonFilterMonitoringClass
	MTRClusterActivatedCarbonFilterMonitoringClassOnce sync.Once
)

func getMTRClusterActivatedCarbonFilterMonitoringClass() _MTRClusterActivatedCarbonFilterMonitoringClass {
	MTRClusterActivatedCarbonFilterMonitoringClassOnce.Do(func() {
		MTRClusterActivatedCarbonFilterMonitoringClass = _MTRClusterActivatedCarbonFilterMonitoringClass{objc.GetClass("MTRClusterActivatedCarbonFilterMonitoring")}
	})
	return MTRClusterActivatedCarbonFilterMonitoringClass
}

type _MTRClusterActivatedCarbonFilterMonitoringClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterActivatedCarbonFilterMonitoring] class.
type IMTRClusterActivatedCarbonFilterMonitoring interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterActivatedCarbonFilterMonitoring
type MTRClusterActivatedCarbonFilterMonitoring struct {
	MTRGenericCluster
}

// MTRClusterActivatedCarbonFilterMonitoringFrom constructs a [MTRClusterActivatedCarbonFilterMonitoring] from an unsafe.Pointer.
func MTRClusterActivatedCarbonFilterMonitoringFrom(ptr unsafe.Pointer) MTRClusterActivatedCarbonFilterMonitoring {
	return MTRClusterActivatedCarbonFilterMonitoring{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterActivatedCarbonFilterMonitoringClass) Alloc() MTRClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRClusterActivatedCarbonFilterMonitoring](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterActivatedCarbonFilterMonitoringClass) New() MTRClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRClusterActivatedCarbonFilterMonitoring](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterActivatedCarbonFilterMonitoring) Init() MTRClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRClusterActivatedCarbonFilterMonitoring](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterActivatedCarbonFilterMonitoring) Autorelease() MTRClusterActivatedCarbonFilterMonitoring {
	rv := objc.Send[MTRClusterActivatedCarbonFilterMonitoring](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterActivatedCarbonFilterMonitoring creates a new MTRClusterActivatedCarbonFilterMonitoring instance.
func NewMTRClusterActivatedCarbonFilterMonitoring() MTRClusterActivatedCarbonFilterMonitoring {
	return getMTRClusterActivatedCarbonFilterMonitoringClass().New()
}
