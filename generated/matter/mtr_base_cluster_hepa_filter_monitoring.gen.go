// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterHEPAFilterMonitoring] class.
var (
	MTRBaseClusterHEPAFilterMonitoringClass     _MTRBaseClusterHEPAFilterMonitoringClass
	MTRBaseClusterHEPAFilterMonitoringClassOnce sync.Once
)

func getMTRBaseClusterHEPAFilterMonitoringClass() _MTRBaseClusterHEPAFilterMonitoringClass {
	MTRBaseClusterHEPAFilterMonitoringClassOnce.Do(func() {
		MTRBaseClusterHEPAFilterMonitoringClass = _MTRBaseClusterHEPAFilterMonitoringClass{objc.GetClass("MTRBaseClusterHEPAFilterMonitoring")}
	})
	return MTRBaseClusterHEPAFilterMonitoringClass
}

type _MTRBaseClusterHEPAFilterMonitoringClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterHEPAFilterMonitoring] class.
type IMTRBaseClusterHEPAFilterMonitoring interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterHEPAFilterMonitoring
type MTRBaseClusterHEPAFilterMonitoring struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterHEPAFilterMonitoringFrom constructs a [MTRBaseClusterHEPAFilterMonitoring] from an unsafe.Pointer.
func MTRBaseClusterHEPAFilterMonitoringFrom(ptr unsafe.Pointer) MTRBaseClusterHEPAFilterMonitoring {
	return MTRBaseClusterHEPAFilterMonitoring{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterHEPAFilterMonitoringClass) Alloc() MTRBaseClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRBaseClusterHEPAFilterMonitoring](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterHEPAFilterMonitoringClass) New() MTRBaseClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRBaseClusterHEPAFilterMonitoring](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterHEPAFilterMonitoring) Init() MTRBaseClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRBaseClusterHEPAFilterMonitoring](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterHEPAFilterMonitoring) Autorelease() MTRBaseClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRBaseClusterHEPAFilterMonitoring](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterHEPAFilterMonitoring creates a new MTRBaseClusterHEPAFilterMonitoring instance.
func NewMTRBaseClusterHEPAFilterMonitoring() MTRBaseClusterHEPAFilterMonitoring {
	return getMTRBaseClusterHEPAFilterMonitoringClass().New()
}




