// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICISession] class.
var (
	mIDICISessionClass     _MIDICISessionClass
	mIDICISessionClassOnce sync.Once
)

func getMIDICISessionClass() _MIDICISessionClass {
	mIDICISessionClassOnce.Do(func() {
		mIDICISessionClass = _MIDICISessionClass{objc.GetClass("MIDICISession")}
	})
	return mIDICISessionClass
}

type _MIDICISessionClass struct {
	class objc.Class
}

// An interface definition for the [MIDICISession] class.
type IMIDICISession interface {
	objectivec.IObject
}

// An object that represents a MIDI-CI session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession
type MIDICISession struct {
	objectivec.Object
}

// MIDICISessionFrom constructs a [MIDICISession] from an unsafe.Pointer.
//
// An object that represents a MIDI-CI session.
func MIDICISessionFrom(ptr unsafe.Pointer) MIDICISession {
	return MIDICISession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICISessionClass) Alloc() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICISessionClass) New() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICISession) Init() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICISession) Autorelease() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICISession creates a new MIDICISession instance.
func NewMIDICISession() MIDICISession {
	return getMIDICISessionClass().New()
}




