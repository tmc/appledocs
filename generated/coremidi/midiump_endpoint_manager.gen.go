// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPEndpointManager] class.
var (
	MIDIUMPEndpointManagerClass     _MIDIUMPEndpointManagerClass
	MIDIUMPEndpointManagerClassOnce sync.Once
)

func getMIDIUMPEndpointManagerClass() _MIDIUMPEndpointManagerClass {
	MIDIUMPEndpointManagerClassOnce.Do(func() {
		MIDIUMPEndpointManagerClass = _MIDIUMPEndpointManagerClass{objc.GetClass("MIDIUMPEndpointManager")}
	})
	return MIDIUMPEndpointManagerClass
}

type _MIDIUMPEndpointManagerClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPEndpointManager] class.
type IMIDIUMPEndpointManager interface {
	objectivec.IObject
	UMPEndpoints() []MIDIUMPEndpoint
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager
type MIDIUMPEndpointManager struct {
	objectivec.Object
}

// MIDIUMPEndpointManagerFrom constructs a [MIDIUMPEndpointManager] from an unsafe.Pointer.
func MIDIUMPEndpointManagerFrom(ptr unsafe.Pointer) MIDIUMPEndpointManager {
	return MIDIUMPEndpointManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPEndpointManagerClass) Alloc() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPEndpointManagerClass) New() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPEndpointManager) Init() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPEndpointManager) Autorelease() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpointManager creates a new MIDIUMPEndpointManager instance.
func NewMIDIUMPEndpointManager() MIDIUMPEndpointManager {
	return getMIDIUMPEndpointManagerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/shared
func (mc _MIDIUMPEndpointManagerClass) SharedInstance() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/shared
func (m_ MIDIUMPEndpointManager) SharedInstance() IMIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m_.ID, objc.Sel("sharedInstance"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/umpEndpoints
func (m_ MIDIUMPEndpointManager) UMPEndpoints() []MIDIUMPEndpoint {
	rv := objc.Send[[]MIDIUMPEndpoint](m_.ID, objc.Sel("UMPEndpoints"))
	return rv
}



