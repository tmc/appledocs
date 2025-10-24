// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZAudioOutputStreamSink] class.
var (
	VZAudioOutputStreamSinkClass     _VZAudioOutputStreamSinkClass
	VZAudioOutputStreamSinkClassOnce sync.Once
)

func getVZAudioOutputStreamSinkClass() _VZAudioOutputStreamSinkClass {
	VZAudioOutputStreamSinkClassOnce.Do(func() {
		VZAudioOutputStreamSinkClass = _VZAudioOutputStreamSinkClass{objc.GetClass("VZAudioOutputStreamSink")}
	})
	return VZAudioOutputStreamSinkClass
}

type _VZAudioOutputStreamSinkClass struct {
	class objc.Class
}

// An interface definition for the [VZAudioOutputStreamSink] class.
type IVZAudioOutputStreamSink interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The base class for an audio output stream sink.
//
// An audio output stream sink defines how the host system consumes audio data from a guest. Don’t instantiate directly, use one of its subclasses, such as instead.


// The base class for an audio output stream sink.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type VZAudioOutputStreamSink struct {
	objectivec.Object
}

// VZAudioOutputStreamSinkFrom constructs a [VZAudioOutputStreamSink] from an unsafe.Pointer.
//
// The base class for an audio output stream sink.
func VZAudioOutputStreamSinkFrom(ptr unsafe.Pointer) VZAudioOutputStreamSink {
	return VZAudioOutputStreamSink{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZAudioOutputStreamSinkClass) Alloc() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZAudioOutputStreamSinkClass) New() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioOutputStreamSink) Init() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioOutputStreamSink) Autorelease() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioOutputStreamSink creates a new VZAudioOutputStreamSink instance.
func NewVZAudioOutputStreamSink() VZAudioOutputStreamSink {
	return getVZAudioOutputStreamSinkClass().New()
}




