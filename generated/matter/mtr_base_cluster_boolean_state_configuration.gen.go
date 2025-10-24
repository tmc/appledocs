// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBooleanStateConfiguration] class.
var (
	MTRBaseClusterBooleanStateConfigurationClass     _MTRBaseClusterBooleanStateConfigurationClass
	MTRBaseClusterBooleanStateConfigurationClassOnce sync.Once
)

func getMTRBaseClusterBooleanStateConfigurationClass() _MTRBaseClusterBooleanStateConfigurationClass {
	MTRBaseClusterBooleanStateConfigurationClassOnce.Do(func() {
		MTRBaseClusterBooleanStateConfigurationClass = _MTRBaseClusterBooleanStateConfigurationClass{objc.GetClass("MTRBaseClusterBooleanStateConfiguration")}
	})
	return MTRBaseClusterBooleanStateConfigurationClass
}

type _MTRBaseClusterBooleanStateConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBooleanStateConfiguration] class.
type IMTRBaseClusterBooleanStateConfiguration interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBooleanStateConfiguration
type MTRBaseClusterBooleanStateConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBooleanStateConfigurationFrom constructs a [MTRBaseClusterBooleanStateConfiguration] from an unsafe.Pointer.
func MTRBaseClusterBooleanStateConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterBooleanStateConfiguration {
	return MTRBaseClusterBooleanStateConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBooleanStateConfigurationClass) Alloc() MTRBaseClusterBooleanStateConfiguration {
	rv := objc.Send[MTRBaseClusterBooleanStateConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBooleanStateConfigurationClass) New() MTRBaseClusterBooleanStateConfiguration {
	rv := objc.Send[MTRBaseClusterBooleanStateConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBooleanStateConfiguration) Init() MTRBaseClusterBooleanStateConfiguration {
	rv := objc.Send[MTRBaseClusterBooleanStateConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBooleanStateConfiguration) Autorelease() MTRBaseClusterBooleanStateConfiguration {
	rv := objc.Send[MTRBaseClusterBooleanStateConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBooleanStateConfiguration creates a new MTRBaseClusterBooleanStateConfiguration instance.
func NewMTRBaseClusterBooleanStateConfiguration() MTRBaseClusterBooleanStateConfiguration {
	return getMTRBaseClusterBooleanStateConfigurationClass().New()
}
