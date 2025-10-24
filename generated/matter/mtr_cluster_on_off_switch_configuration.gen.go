// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOnOffSwitchConfiguration] class.
var (
	MTRClusterOnOffSwitchConfigurationClass     _MTRClusterOnOffSwitchConfigurationClass
	MTRClusterOnOffSwitchConfigurationClassOnce sync.Once
)

func getMTRClusterOnOffSwitchConfigurationClass() _MTRClusterOnOffSwitchConfigurationClass {
	MTRClusterOnOffSwitchConfigurationClassOnce.Do(func() {
		MTRClusterOnOffSwitchConfigurationClass = _MTRClusterOnOffSwitchConfigurationClass{objc.GetClass("MTRClusterOnOffSwitchConfiguration")}
	})
	return MTRClusterOnOffSwitchConfigurationClass
}

type _MTRClusterOnOffSwitchConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOnOffSwitchConfiguration] class.
type IMTRClusterOnOffSwitchConfiguration interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOnOffSwitchConfiguration
type MTRClusterOnOffSwitchConfiguration struct {
	MTRGenericCluster
}

// MTRClusterOnOffSwitchConfigurationFrom constructs a [MTRClusterOnOffSwitchConfiguration] from an unsafe.Pointer.
func MTRClusterOnOffSwitchConfigurationFrom(ptr unsafe.Pointer) MTRClusterOnOffSwitchConfiguration {
	return MTRClusterOnOffSwitchConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOnOffSwitchConfigurationClass) Alloc() MTRClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRClusterOnOffSwitchConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOnOffSwitchConfigurationClass) New() MTRClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRClusterOnOffSwitchConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOnOffSwitchConfiguration) Init() MTRClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRClusterOnOffSwitchConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOnOffSwitchConfiguration) Autorelease() MTRClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRClusterOnOffSwitchConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOnOffSwitchConfiguration creates a new MTRClusterOnOffSwitchConfiguration instance.
func NewMTRClusterOnOffSwitchConfiguration() MTRClusterOnOffSwitchConfiguration {
	return getMTRClusterOnOffSwitchConfigurationClass().New()
}




