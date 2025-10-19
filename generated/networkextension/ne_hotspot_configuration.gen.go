// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotConfiguration] class.
var (
	nEHotspotConfigurationClass     _NEHotspotConfigurationClass
	nEHotspotConfigurationClassOnce sync.Once
)

func getNEHotspotConfigurationClass() _NEHotspotConfigurationClass {
	nEHotspotConfigurationClassOnce.Do(func() {
		nEHotspotConfigurationClass = _NEHotspotConfigurationClass{objc.GetClass("NEHotspotConfiguration")}
	})
	return nEHotspotConfigurationClass
}

type _NEHotspotConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotConfiguration] class.
type INEHotspotConfiguration interface {
	objectivec.IObject
}

// Configuration settings for a Wi-Fi network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration
type NEHotspotConfiguration struct {
	objectivec.Object
}

// NEHotspotConfigurationFrom constructs a [NEHotspotConfiguration] from an unsafe.Pointer.
//
// Configuration settings for a Wi-Fi network.
func NEHotspotConfigurationFrom(ptr unsafe.Pointer) NEHotspotConfiguration {
	return NEHotspotConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotConfigurationClass) Alloc() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotConfigurationClass) New() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotConfiguration) Init() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotConfiguration) Autorelease() NEHotspotConfiguration {
	rv := objc.Send[NEHotspotConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotConfiguration creates a new NEHotspotConfiguration instance.
func NewNEHotspotConfiguration() NEHotspotConfiguration {
	return getNEHotspotConfigurationClass().New()
}




