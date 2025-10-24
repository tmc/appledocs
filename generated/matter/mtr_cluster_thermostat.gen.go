// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterThermostat] class.
var (
	MTRClusterThermostatClass     _MTRClusterThermostatClass
	MTRClusterThermostatClassOnce sync.Once
)

func getMTRClusterThermostatClass() _MTRClusterThermostatClass {
	MTRClusterThermostatClassOnce.Do(func() {
		MTRClusterThermostatClass = _MTRClusterThermostatClass{objc.GetClass("MTRClusterThermostat")}
	})
	return MTRClusterThermostatClass
}

type _MTRClusterThermostatClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterThermostat] class.
type IMTRClusterThermostat interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThermostat
type MTRClusterThermostat struct {
	MTRGenericCluster
}

// MTRClusterThermostatFrom constructs a [MTRClusterThermostat] from an unsafe.Pointer.
func MTRClusterThermostatFrom(ptr unsafe.Pointer) MTRClusterThermostat {
	return MTRClusterThermostat{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThermostatClass) Alloc() MTRClusterThermostat {
	rv := objc.Send[MTRClusterThermostat](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterThermostatClass) New() MTRClusterThermostat {
	rv := objc.Send[MTRClusterThermostat](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThermostat) Init() MTRClusterThermostat {
	rv := objc.Send[MTRClusterThermostat](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThermostat) Autorelease() MTRClusterThermostat {
	rv := objc.Send[MTRClusterThermostat](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThermostat creates a new MTRClusterThermostat instance.
func NewMTRClusterThermostat() MTRClusterThermostat {
	return getMTRClusterThermostatClass().New()
}
