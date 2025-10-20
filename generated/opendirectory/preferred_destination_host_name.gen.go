// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [preferredDestinationHostName] class.
var (
	PreferredDestinationHostNameClass     _preferredDestinationHostNameClass
	PreferredDestinationHostNameClassOnce sync.Once
)

func getpreferredDestinationHostNameClass() _preferredDestinationHostNameClass {
	PreferredDestinationHostNameClassOnce.Do(func() {
		PreferredDestinationHostNameClass = _preferredDestinationHostNameClass{objc.GetClass("preferredDestinationHostName")}
	})
	return PreferredDestinationHostNameClass
}

type _preferredDestinationHostNameClass struct {
	class objc.Class
}

// An interface definition for the [preferredDestinationHostName] class.
type IpreferredDestinationHostName interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-c.ivar
type preferredDestinationHostName struct {
	objectivec.Object
}

// preferredDestinationHostNameFrom constructs a [preferredDestinationHostName] from an unsafe.Pointer.
func preferredDestinationHostNameFrom(ptr unsafe.Pointer) preferredDestinationHostName {
	return preferredDestinationHostName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _preferredDestinationHostNameClass) Alloc() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _preferredDestinationHostNameClass) New() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ preferredDestinationHostName) Init() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ preferredDestinationHostName) Autorelease() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpreferredDestinationHostName creates a new preferredDestinationHostName instance.
func NewpreferredDestinationHostName() preferredDestinationHostName {
	return getpreferredDestinationHostNameClass().New()
}




