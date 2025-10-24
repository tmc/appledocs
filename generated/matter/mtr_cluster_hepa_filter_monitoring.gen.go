// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterHEPAFilterMonitoring] class.
var (
	MTRClusterHEPAFilterMonitoringClass     _MTRClusterHEPAFilterMonitoringClass
	MTRClusterHEPAFilterMonitoringClassOnce sync.Once
)

func getMTRClusterHEPAFilterMonitoringClass() _MTRClusterHEPAFilterMonitoringClass {
	MTRClusterHEPAFilterMonitoringClassOnce.Do(func() {
		MTRClusterHEPAFilterMonitoringClass = _MTRClusterHEPAFilterMonitoringClass{objc.GetClass("MTRClusterHEPAFilterMonitoring")}
	})
	return MTRClusterHEPAFilterMonitoringClass
}

type _MTRClusterHEPAFilterMonitoringClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterHEPAFilterMonitoring] class.
type IMTRClusterHEPAFilterMonitoring interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterHEPAFilterMonitoring
type MTRClusterHEPAFilterMonitoring struct {
	MTRGenericCluster
}

// MTRClusterHEPAFilterMonitoringFrom constructs a [MTRClusterHEPAFilterMonitoring] from an unsafe.Pointer.
func MTRClusterHEPAFilterMonitoringFrom(ptr unsafe.Pointer) MTRClusterHEPAFilterMonitoring {
	return MTRClusterHEPAFilterMonitoring{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterHEPAFilterMonitoringClass) Alloc() MTRClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRClusterHEPAFilterMonitoring](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterHEPAFilterMonitoringClass) New() MTRClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRClusterHEPAFilterMonitoring](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterHEPAFilterMonitoring) Init() MTRClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRClusterHEPAFilterMonitoring](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterHEPAFilterMonitoring) Autorelease() MTRClusterHEPAFilterMonitoring {
	rv := objc.Send[MTRClusterHEPAFilterMonitoring](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterHEPAFilterMonitoring creates a new MTRClusterHEPAFilterMonitoring instance.
func NewMTRClusterHEPAFilterMonitoring() MTRClusterHEPAFilterMonitoring {
	return getMTRClusterHEPAFilterMonitoringClass().New()
}




