// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hideRegistration] class.
var (
	HideRegistrationClass     _hideRegistrationClass
	HideRegistrationClassOnce sync.Once
)

func gethideRegistrationClass() _hideRegistrationClass {
	HideRegistrationClassOnce.Do(func() {
		HideRegistrationClass = _hideRegistrationClass{objc.GetClass("hideRegistration")}
	})
	return HideRegistrationClass
}

type _hideRegistrationClass struct {
	class objc.Class
}

// An interface definition for the [hideRegistration] class.
type IhideRegistration interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-c.ivar
type hideRegistration struct {
	objectivec.Object
}

// hideRegistrationFrom constructs a [hideRegistration] from an unsafe.Pointer.
func hideRegistrationFrom(ptr unsafe.Pointer) hideRegistration {
	return hideRegistration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _hideRegistrationClass) Alloc() hideRegistration {
	rv := objc.Send[hideRegistration](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _hideRegistrationClass) New() hideRegistration {
	rv := objc.Send[hideRegistration](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hideRegistration) Init() hideRegistration {
	rv := objc.Send[hideRegistration](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hideRegistration) Autorelease() hideRegistration {
	rv := objc.Send[hideRegistration](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhideRegistration creates a new hideRegistration instance.
func NewhideRegistration() hideRegistration {
	return gethideRegistrationClass().New()
}




