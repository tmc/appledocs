// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBooleanStateConfiguration] class.
var (
	MTRClusterBooleanStateConfigurationClass     _MTRClusterBooleanStateConfigurationClass
	MTRClusterBooleanStateConfigurationClassOnce sync.Once
)

func getMTRClusterBooleanStateConfigurationClass() _MTRClusterBooleanStateConfigurationClass {
	MTRClusterBooleanStateConfigurationClassOnce.Do(func() {
		MTRClusterBooleanStateConfigurationClass = _MTRClusterBooleanStateConfigurationClass{objc.GetClass("MTRClusterBooleanStateConfiguration")}
	})
	return MTRClusterBooleanStateConfigurationClass
}

type _MTRClusterBooleanStateConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBooleanStateConfiguration] class.
type IMTRClusterBooleanStateConfiguration interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBooleanStateConfiguration
type MTRClusterBooleanStateConfiguration struct {
	MTRGenericCluster
}

// MTRClusterBooleanStateConfigurationFrom constructs a [MTRClusterBooleanStateConfiguration] from an unsafe.Pointer.
func MTRClusterBooleanStateConfigurationFrom(ptr unsafe.Pointer) MTRClusterBooleanStateConfiguration {
	return MTRClusterBooleanStateConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBooleanStateConfigurationClass) Alloc() MTRClusterBooleanStateConfiguration {
	rv := objc.Send[MTRClusterBooleanStateConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBooleanStateConfigurationClass) New() MTRClusterBooleanStateConfiguration {
	rv := objc.Send[MTRClusterBooleanStateConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBooleanStateConfiguration) Init() MTRClusterBooleanStateConfiguration {
	rv := objc.Send[MTRClusterBooleanStateConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBooleanStateConfiguration) Autorelease() MTRClusterBooleanStateConfiguration {
	rv := objc.Send[MTRClusterBooleanStateConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBooleanStateConfiguration creates a new MTRClusterBooleanStateConfiguration instance.
func NewMTRClusterBooleanStateConfiguration() MTRClusterBooleanStateConfiguration {
	return getMTRClusterBooleanStateConfigurationClass().New()
}
