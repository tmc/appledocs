// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterProviderConfiguration] class.
var nEFilterProviderConfigurationClass = _NEFilterProviderConfigurationClass{objc.GetClass("NEFilterProviderConfiguration")}

type _NEFilterProviderConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterProviderConfiguration] class.
type INEFilterProviderConfiguration interface {
	objectivec.IObject
}

// Configuration parameters for a content filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration

type NEFilterProviderConfiguration struct {
	objectivec.Object
}

// NEFilterProviderConfigurationFrom constructs a [NEFilterProviderConfiguration] from an unsafe.Pointer.
//
// Configuration parameters for a content filter.
func NEFilterProviderConfigurationFrom(ptr unsafe.Pointer) NEFilterProviderConfiguration {
	return NEFilterProviderConfiguration{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NEFilterProviderConfigurationClass) Alloc() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NEFilterProviderConfigurationClass) New() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterProviderConfiguration) Init() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterProviderConfiguration) Autorelease() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterProviderConfiguration creates a new NEFilterProviderConfiguration instance.
func NewNEFilterProviderConfiguration() NEFilterProviderConfiguration {
	return nEFilterProviderConfigurationClass.New()
}




