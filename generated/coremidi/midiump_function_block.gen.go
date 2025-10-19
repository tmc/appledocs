// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPFunctionBlock] class.
var (
	mIDIUMPFunctionBlockClass     _MIDIUMPFunctionBlockClass
	mIDIUMPFunctionBlockClassOnce sync.Once
)

func getMIDIUMPFunctionBlockClass() _MIDIUMPFunctionBlockClass {
	mIDIUMPFunctionBlockClassOnce.Do(func() {
		mIDIUMPFunctionBlockClass = _MIDIUMPFunctionBlockClass{objc.GetClass("MIDIUMPFunctionBlock")}
	})
	return mIDIUMPFunctionBlockClass
}

type _MIDIUMPFunctionBlockClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPFunctionBlock] class.
type IMIDIUMPFunctionBlock interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock
type MIDIUMPFunctionBlock struct {
	objectivec.Object
}

// MIDIUMPFunctionBlockFrom constructs a [MIDIUMPFunctionBlock] from an unsafe.Pointer.
func MIDIUMPFunctionBlockFrom(ptr unsafe.Pointer) MIDIUMPFunctionBlock {
	return MIDIUMPFunctionBlock{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPFunctionBlockClass) Alloc() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPFunctionBlockClass) New() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPFunctionBlock) Init() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPFunctionBlock) Autorelease() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPFunctionBlock creates a new MIDIUMPFunctionBlock instance.
func NewMIDIUMPFunctionBlock() MIDIUMPFunctionBlock {
	return getMIDIUMPFunctionBlockClass().New()
}




