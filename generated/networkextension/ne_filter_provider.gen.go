// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterProvider] class.
var (
	NEFilterProviderClass     _NEFilterProviderClass
	NEFilterProviderClassOnce sync.Once
)

func getNEFilterProviderClass() _NEFilterProviderClass {
	NEFilterProviderClassOnce.Do(func() {
		NEFilterProviderClass = _NEFilterProviderClass{objc.GetClass("NEFilterProvider")}
	})
	return NEFilterProviderClass
}

type _NEFilterProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterProvider] class.
type INEFilterProvider interface {
	INEProvider
}

// An abstract base class shared by content filters.
//
// A Network Content Filter is made up of two Filter Provider extensions: The examines network content as it passes through the network stack on the device and decides if the network content should be blocked or allowed to pass on to its final destination. Because the Filter Data Provider extension has access to all of the network content flowing through the device, it runs in a very restrictive sandbox. The sandbox prevents the Filter Data Provider extension from moving network content outside of its address space by blocking all network access, IPC, and disk write operations. The Filter Data Provider extension is implemented by creating a custom subclass of the class. The is responsible for feeding information to the Filter Data Provider extension so that the Filter Data Provider extension can do its job. For example, the Filter Control Provider extension can be notified by the Filter Data Provider extension that it does not have enough information to make a decision about a particular flow of network content. The Filter Control Provider extension can then download more filtering rules from a server and write the rules to a location where the Filter Data Provider can access them. The Filter Control Provider extension is implemented by creating a custom subclass of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider
type NEFilterProvider struct {
	NEProvider
}

// NEFilterProviderFrom constructs a [NEFilterProvider] from an unsafe.Pointer.
//
// An abstract base class shared by content filters.
func NEFilterProviderFrom(ptr unsafe.Pointer) NEFilterProvider {
	return NEFilterProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterProviderClass) Alloc() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterProviderClass) New() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterProvider) Init() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterProvider) Autorelease() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterProvider creates a new NEFilterProvider instance.
func NewNEFilterProvider() NEFilterProvider {
	return getNEFilterProviderClass().New()
}


// The domain for errors resulting from calls to the filter manager.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (n_ NEFilterProvider) NEFilterErrorDomain() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterErrorDomain"))
	return rv
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterprovider/filterconfiguration
func (n_ NEFilterProvider) FilterConfiguration() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("filterConfiguration"))
	return rv
}


// SetFilterConfiguration sets the value of the filterConfiguration property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterprovider/filterconfiguration
func (n_ NEFilterProvider) SetFilterConfiguration(value INEFilterProviderConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterConfiguration:"), value)
}



