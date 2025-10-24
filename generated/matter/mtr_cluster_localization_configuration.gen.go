// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterLocalizationConfiguration] class.
var (
	MTRClusterLocalizationConfigurationClass     _MTRClusterLocalizationConfigurationClass
	MTRClusterLocalizationConfigurationClassOnce sync.Once
)

func getMTRClusterLocalizationConfigurationClass() _MTRClusterLocalizationConfigurationClass {
	MTRClusterLocalizationConfigurationClassOnce.Do(func() {
		MTRClusterLocalizationConfigurationClass = _MTRClusterLocalizationConfigurationClass{objc.GetClass("MTRClusterLocalizationConfiguration")}
	})
	return MTRClusterLocalizationConfigurationClass
}

type _MTRClusterLocalizationConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLocalizationConfiguration] class.
type IMTRClusterLocalizationConfiguration interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLocalizationConfiguration
type MTRClusterLocalizationConfiguration struct {
	MTRGenericCluster
}

// MTRClusterLocalizationConfigurationFrom constructs a [MTRClusterLocalizationConfiguration] from an unsafe.Pointer.
func MTRClusterLocalizationConfigurationFrom(ptr unsafe.Pointer) MTRClusterLocalizationConfiguration {
	return MTRClusterLocalizationConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLocalizationConfigurationClass) Alloc() MTRClusterLocalizationConfiguration {
	rv := objc.Send[MTRClusterLocalizationConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLocalizationConfigurationClass) New() MTRClusterLocalizationConfiguration {
	rv := objc.Send[MTRClusterLocalizationConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLocalizationConfiguration) Init() MTRClusterLocalizationConfiguration {
	rv := objc.Send[MTRClusterLocalizationConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLocalizationConfiguration) Autorelease() MTRClusterLocalizationConfiguration {
	rv := objc.Send[MTRClusterLocalizationConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLocalizationConfiguration creates a new MTRClusterLocalizationConfiguration instance.
func NewMTRClusterLocalizationConfiguration() MTRClusterLocalizationConfiguration {
	return getMTRClusterLocalizationConfigurationClass().New()
}




