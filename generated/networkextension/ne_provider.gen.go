// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEProvider] class.
var (
	NEProviderClass     _NEProviderClass
	NEProviderClassOnce sync.Once
)

func getNEProviderClass() _NEProviderClass {
	NEProviderClassOnce.Do(func() {
		NEProviderClass = _NEProviderClass{objc.GetClass("NEProvider")}
	})
	return NEProviderClass
}

type _NEProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEProvider] class.
type INEProvider interface {
	objectivec.IObject
	// properties:
	DefaultPath() objc.IObject /* cross-framework: NWPath */
	SetDefaultPath(value objc.IObject /* cross-framework: NWPath */)
	// methods:
}

// An abstract base class for all NetworkExtension providers.
//
// See the documentation for the subclasses for details about how to create Network Extension Provider extensions. The class and its subclasses expose methods and properties that allow Network Extension Provider extensions to participate in and affect the network data path on iOS and macOS. For example, the method in allows Filter Data Provider extensions to make pass/block decisions on TCP connections as the connections are established on the system.


// An abstract base class for all NetworkExtension providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider
type NEProvider struct {
	objectivec.Object
}

// NEProviderFrom constructs a [NEProvider] from an unsafe.Pointer.
//
// An abstract base class for all NetworkExtension providers.
func NEProviderFrom(ptr unsafe.Pointer) NEProvider {
	return NEProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEProviderClass) Alloc() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEProviderClass) New() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProvider) Init() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProvider) Autorelease() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProvider creates a new NEProvider instance.
func NewNEProvider() NEProvider {
	return getNEProviderClass().New()
}



// The current default network path used for connections created by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprovider/defaultpath
func (n_ NEProvider) DefaultPath() objc.IObject /* cross-framework: NWPath */ {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("defaultPath"))
	return rv
}


// The current default network path used for connections created by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neprovider/defaultpath
func (n_ NEProvider) SetDefaultPath(value objc.IObject /* cross-framework: NWPath */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDefaultPath:"), value)
}



