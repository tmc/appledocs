// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPowerSourceConfiguration] class.
var (
	MTRClusterPowerSourceConfigurationClass     _MTRClusterPowerSourceConfigurationClass
	MTRClusterPowerSourceConfigurationClassOnce sync.Once
)

func getMTRClusterPowerSourceConfigurationClass() _MTRClusterPowerSourceConfigurationClass {
	MTRClusterPowerSourceConfigurationClassOnce.Do(func() {
		MTRClusterPowerSourceConfigurationClass = _MTRClusterPowerSourceConfigurationClass{objc.GetClass("MTRClusterPowerSourceConfiguration")}
	})
	return MTRClusterPowerSourceConfigurationClass
}

type _MTRClusterPowerSourceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPowerSourceConfiguration] class.
type IMTRClusterPowerSourceConfiguration interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerSourceConfiguration
type MTRClusterPowerSourceConfiguration struct {
	MTRGenericCluster
}

// MTRClusterPowerSourceConfigurationFrom constructs a [MTRClusterPowerSourceConfiguration] from an unsafe.Pointer.
func MTRClusterPowerSourceConfigurationFrom(ptr unsafe.Pointer) MTRClusterPowerSourceConfiguration {
	return MTRClusterPowerSourceConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPowerSourceConfigurationClass) Alloc() MTRClusterPowerSourceConfiguration {
	rv := objc.Send[MTRClusterPowerSourceConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPowerSourceConfigurationClass) New() MTRClusterPowerSourceConfiguration {
	rv := objc.Send[MTRClusterPowerSourceConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPowerSourceConfiguration) Init() MTRClusterPowerSourceConfiguration {
	rv := objc.Send[MTRClusterPowerSourceConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPowerSourceConfiguration) Autorelease() MTRClusterPowerSourceConfiguration {
	rv := objc.Send[MTRClusterPowerSourceConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPowerSourceConfiguration creates a new MTRClusterPowerSourceConfiguration instance.
func NewMTRClusterPowerSourceConfiguration() MTRClusterPowerSourceConfiguration {
	return getMTRClusterPowerSourceConfigurationClass().New()
}




