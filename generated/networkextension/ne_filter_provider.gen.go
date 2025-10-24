// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other NetworkExtension classes.


// A parent class referenced by other NetworkExtension classes. [Full Topic]
type NEFilterProvider struct {
	objectivec.Object
}

// NEFilterProviderFrom constructs a [NEFilterProvider] from an unsafe.Pointer.
//
// A parent class referenced by other NetworkExtension classes.
func NEFilterProviderFrom(ptr unsafe.Pointer) NEFilterProvider {
	return NEFilterProvider{objectivec.Object{objc.ID(ptr)}}
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




