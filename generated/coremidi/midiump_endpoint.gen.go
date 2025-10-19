// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPEndpoint] class.
var (
	mIDIUMPEndpointClass     _MIDIUMPEndpointClass
	mIDIUMPEndpointClassOnce sync.Once
)

func getMIDIUMPEndpointClass() _MIDIUMPEndpointClass {
	mIDIUMPEndpointClassOnce.Do(func() {
		mIDIUMPEndpointClass = _MIDIUMPEndpointClass{objc.GetClass("MIDIUMPEndpoint")}
	})
	return mIDIUMPEndpointClass
}

type _MIDIUMPEndpointClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPEndpoint] class.
type IMIDIUMPEndpoint interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint
type MIDIUMPEndpoint struct {
	objectivec.Object
}

// MIDIUMPEndpointFrom constructs a [MIDIUMPEndpoint] from an unsafe.Pointer.
func MIDIUMPEndpointFrom(ptr unsafe.Pointer) MIDIUMPEndpoint {
	return MIDIUMPEndpoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPEndpointClass) Alloc() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPEndpointClass) New() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPEndpoint) Init() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPEndpoint) Autorelease() MIDIUMPEndpoint {
	rv := objc.Send[MIDIUMPEndpoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpoint creates a new MIDIUMPEndpoint instance.
func NewMIDIUMPEndpoint() MIDIUMPEndpoint {
	return getMIDIUMPEndpointClass().New()
}




