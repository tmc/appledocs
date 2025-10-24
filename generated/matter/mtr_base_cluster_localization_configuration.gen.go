// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterLocalizationConfiguration] class.
var (
	MTRBaseClusterLocalizationConfigurationClass     _MTRBaseClusterLocalizationConfigurationClass
	MTRBaseClusterLocalizationConfigurationClassOnce sync.Once
)

func getMTRBaseClusterLocalizationConfigurationClass() _MTRBaseClusterLocalizationConfigurationClass {
	MTRBaseClusterLocalizationConfigurationClassOnce.Do(func() {
		MTRBaseClusterLocalizationConfigurationClass = _MTRBaseClusterLocalizationConfigurationClass{objc.GetClass("MTRBaseClusterLocalizationConfiguration")}
	})
	return MTRBaseClusterLocalizationConfigurationClass
}

type _MTRBaseClusterLocalizationConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterLocalizationConfiguration] class.
type IMTRBaseClusterLocalizationConfiguration interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLocalizationConfiguration
type MTRBaseClusterLocalizationConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLocalizationConfigurationFrom constructs a [MTRBaseClusterLocalizationConfiguration] from an unsafe.Pointer.
func MTRBaseClusterLocalizationConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterLocalizationConfiguration {
	return MTRBaseClusterLocalizationConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLocalizationConfigurationClass) Alloc() MTRBaseClusterLocalizationConfiguration {
	rv := objc.Send[MTRBaseClusterLocalizationConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterLocalizationConfigurationClass) New() MTRBaseClusterLocalizationConfiguration {
	rv := objc.Send[MTRBaseClusterLocalizationConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLocalizationConfiguration) Init() MTRBaseClusterLocalizationConfiguration {
	rv := objc.Send[MTRBaseClusterLocalizationConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLocalizationConfiguration) Autorelease() MTRBaseClusterLocalizationConfiguration {
	rv := objc.Send[MTRBaseClusterLocalizationConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLocalizationConfiguration creates a new MTRBaseClusterLocalizationConfiguration instance.
func NewMTRBaseClusterLocalizationConfiguration() MTRBaseClusterLocalizationConfiguration {
	return getMTRBaseClusterLocalizationConfigurationClass().New()
}
