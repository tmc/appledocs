// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZHostAudioOutputStreamSink] class.
var (
	VZHostAudioOutputStreamSinkClass     _VZHostAudioOutputStreamSinkClass
	VZHostAudioOutputStreamSinkClassOnce sync.Once
)

func getVZHostAudioOutputStreamSinkClass() _VZHostAudioOutputStreamSinkClass {
	VZHostAudioOutputStreamSinkClassOnce.Do(func() {
		VZHostAudioOutputStreamSinkClass = _VZHostAudioOutputStreamSinkClass{objc.GetClass("VZHostAudioOutputStreamSink")}
	})
	return VZHostAudioOutputStreamSinkClass
}

type _VZHostAudioOutputStreamSinkClass struct {
	class objc.Class
}

// An interface definition for the [VZHostAudioOutputStreamSink] class.
type IVZHostAudioOutputStreamSink interface {
	IVZAudioOutputStreamSink
	// properties:
	// methods:
}

// Host audio output stream sink plays audio to the host system’s default output device.
//
// Host output data goes to the same device that uses.


// Host audio output stream sink plays audio to the host system’s default output device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZHostAudioOutputStreamSink
type VZHostAudioOutputStreamSink struct {
	VZAudioOutputStreamSink
}

// VZHostAudioOutputStreamSinkFrom constructs a [VZHostAudioOutputStreamSink] from an unsafe.Pointer.
//
// Host audio output stream sink plays audio to the host system’s default output device.
func VZHostAudioOutputStreamSinkFrom(ptr unsafe.Pointer) VZHostAudioOutputStreamSink {
	return VZHostAudioOutputStreamSink{
		VZAudioOutputStreamSink: VZAudioOutputStreamSinkFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZHostAudioOutputStreamSinkClass) Alloc() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZHostAudioOutputStreamSinkClass) New() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZHostAudioOutputStreamSink) Init() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZHostAudioOutputStreamSink) Autorelease() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioOutputStreamSink creates a new VZHostAudioOutputStreamSink instance.
func NewVZHostAudioOutputStreamSink() VZHostAudioOutputStreamSink {
	return getVZHostAudioOutputStreamSinkClass().New()
}




