// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [preferredDestinationHostPort] class.
var (
	PreferredDestinationHostPortClass     _preferredDestinationHostPortClass
	PreferredDestinationHostPortClassOnce sync.Once
)

func getpreferredDestinationHostPortClass() _preferredDestinationHostPortClass {
	PreferredDestinationHostPortClassOnce.Do(func() {
		PreferredDestinationHostPortClass = _preferredDestinationHostPortClass{objc.GetClass("preferredDestinationHostPort")}
	})
	return PreferredDestinationHostPortClass
}

type _preferredDestinationHostPortClass struct {
	class objc.Class
}

// An interface definition for the [preferredDestinationHostPort] class.
type IpreferredDestinationHostPort interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-c.ivar
type preferredDestinationHostPort struct {
	objectivec.Object
}

// preferredDestinationHostPortFrom constructs a [preferredDestinationHostPort] from an unsafe.Pointer.
func preferredDestinationHostPortFrom(ptr unsafe.Pointer) preferredDestinationHostPort {
	return preferredDestinationHostPort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _preferredDestinationHostPortClass) Alloc() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _preferredDestinationHostPortClass) New() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ preferredDestinationHostPort) Init() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ preferredDestinationHostPort) Autorelease() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpreferredDestinationHostPort creates a new preferredDestinationHostPort instance.
func NewpreferredDestinationHostPort() preferredDestinationHostPort {
	return getpreferredDestinationHostPortClass().New()
}




