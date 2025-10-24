// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterThermostat] class.
var (
	MTRBaseClusterThermostatClass     _MTRBaseClusterThermostatClass
	MTRBaseClusterThermostatClassOnce sync.Once
)

func getMTRBaseClusterThermostatClass() _MTRBaseClusterThermostatClass {
	MTRBaseClusterThermostatClassOnce.Do(func() {
		MTRBaseClusterThermostatClass = _MTRBaseClusterThermostatClass{objc.GetClass("MTRBaseClusterThermostat")}
	})
	return MTRBaseClusterThermostatClass
}

type _MTRBaseClusterThermostatClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterThermostat] class.
type IMTRBaseClusterThermostat interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThermostat
type MTRBaseClusterThermostat struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThermostatFrom constructs a [MTRBaseClusterThermostat] from an unsafe.Pointer.
func MTRBaseClusterThermostatFrom(ptr unsafe.Pointer) MTRBaseClusterThermostat {
	return MTRBaseClusterThermostat{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThermostatClass) Alloc() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterThermostatClass) New() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThermostat) Init() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThermostat) Autorelease() MTRBaseClusterThermostat {
	rv := objc.Send[MTRBaseClusterThermostat](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThermostat creates a new MTRBaseClusterThermostat instance.
func NewMTRBaseClusterThermostat() MTRBaseClusterThermostat {
	return getMTRBaseClusterThermostatClass().New()
}




