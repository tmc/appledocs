// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOnOffSwitchConfiguration] class.
var (
	MTRBaseClusterOnOffSwitchConfigurationClass     _MTRBaseClusterOnOffSwitchConfigurationClass
	MTRBaseClusterOnOffSwitchConfigurationClassOnce sync.Once
)

func getMTRBaseClusterOnOffSwitchConfigurationClass() _MTRBaseClusterOnOffSwitchConfigurationClass {
	MTRBaseClusterOnOffSwitchConfigurationClassOnce.Do(func() {
		MTRBaseClusterOnOffSwitchConfigurationClass = _MTRBaseClusterOnOffSwitchConfigurationClass{objc.GetClass("MTRBaseClusterOnOffSwitchConfiguration")}
	})
	return MTRBaseClusterOnOffSwitchConfigurationClass
}

type _MTRBaseClusterOnOffSwitchConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOnOffSwitchConfiguration] class.
type IMTRBaseClusterOnOffSwitchConfiguration interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOnOffSwitchConfiguration
type MTRBaseClusterOnOffSwitchConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOnOffSwitchConfigurationFrom constructs a [MTRBaseClusterOnOffSwitchConfiguration] from an unsafe.Pointer.
func MTRBaseClusterOnOffSwitchConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterOnOffSwitchConfiguration {
	return MTRBaseClusterOnOffSwitchConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOnOffSwitchConfigurationClass) Alloc() MTRBaseClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRBaseClusterOnOffSwitchConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOnOffSwitchConfigurationClass) New() MTRBaseClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRBaseClusterOnOffSwitchConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOnOffSwitchConfiguration) Init() MTRBaseClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRBaseClusterOnOffSwitchConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOnOffSwitchConfiguration) Autorelease() MTRBaseClusterOnOffSwitchConfiguration {
	rv := objc.Send[MTRBaseClusterOnOffSwitchConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOnOffSwitchConfiguration creates a new MTRBaseClusterOnOffSwitchConfiguration instance.
func NewMTRBaseClusterOnOffSwitchConfiguration() MTRBaseClusterOnOffSwitchConfiguration {
	return getMTRBaseClusterOnOffSwitchConfigurationClass().New()
}




