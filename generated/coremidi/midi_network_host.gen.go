// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDINetworkHost] class.
var (
	mIDINetworkHostClass     _MIDINetworkHostClass
	mIDINetworkHostClassOnce sync.Once
)

func getMIDINetworkHostClass() _MIDINetworkHostClass {
	mIDINetworkHostClassOnce.Do(func() {
		mIDINetworkHostClass = _MIDINetworkHostClass{objc.GetClass("MIDINetworkHost")}
	})
	return mIDINetworkHostClass
}

type _MIDINetworkHostClass struct {
	class objc.Class
}

// An interface definition for the [MIDINetworkHost] class.
type IMIDINetworkHost interface {
	objectivec.IObject
}

// An object that represents the host’s network address.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost
type MIDINetworkHost struct {
	objectivec.Object
}

// MIDINetworkHostFrom constructs a [MIDINetworkHost] from an unsafe.Pointer.
//
// An object that represents the host’s network address.
func MIDINetworkHostFrom(ptr unsafe.Pointer) MIDINetworkHost {
	return MIDINetworkHost{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkHostClass) Alloc() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDINetworkHostClass) New() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkHost) Init() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkHost) Autorelease() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkHost creates a new MIDINetworkHost instance.
func NewMIDINetworkHost() MIDINetworkHost {
	return getMIDINetworkHostClass().New()
}




