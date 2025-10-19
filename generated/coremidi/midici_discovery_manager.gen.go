// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDiscoveryManager] class.
var (
	mIDICIDiscoveryManagerClass     _MIDICIDiscoveryManagerClass
	mIDICIDiscoveryManagerClassOnce sync.Once
)

func getMIDICIDiscoveryManagerClass() _MIDICIDiscoveryManagerClass {
	mIDICIDiscoveryManagerClassOnce.Do(func() {
		mIDICIDiscoveryManagerClass = _MIDICIDiscoveryManagerClass{objc.GetClass("MIDICIDiscoveryManager")}
	})
	return mIDICIDiscoveryManagerClass
}

type _MIDICIDiscoveryManagerClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDiscoveryManager] class.
type IMIDICIDiscoveryManager interface {
	objectivec.IObject
}

// A singleton object that performs systemwide MIDI-CI discovery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager
type MIDICIDiscoveryManager struct {
	objectivec.Object
}

// MIDICIDiscoveryManagerFrom constructs a [MIDICIDiscoveryManager] from an unsafe.Pointer.
//
// A singleton object that performs systemwide MIDI-CI discovery.
func MIDICIDiscoveryManagerFrom(ptr unsafe.Pointer) MIDICIDiscoveryManager {
	return MIDICIDiscoveryManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIDiscoveryManagerClass) Alloc() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIDiscoveryManagerClass) New() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDiscoveryManager) Init() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDiscoveryManager) Autorelease() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveryManager creates a new MIDICIDiscoveryManager instance.
func NewMIDICIDiscoveryManager() MIDICIDiscoveryManager {
	return getMIDICIDiscoveryManagerClass().New()
}




