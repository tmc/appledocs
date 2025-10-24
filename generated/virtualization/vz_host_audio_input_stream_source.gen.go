// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZHostAudioInputStreamSource] class.
var (
	VZHostAudioInputStreamSourceClass     _VZHostAudioInputStreamSourceClass
	VZHostAudioInputStreamSourceClassOnce sync.Once
)

func getVZHostAudioInputStreamSourceClass() _VZHostAudioInputStreamSourceClass {
	VZHostAudioInputStreamSourceClassOnce.Do(func() {
		VZHostAudioInputStreamSourceClass = _VZHostAudioInputStreamSourceClass{objc.GetClass("VZHostAudioInputStreamSource")}
	})
	return VZHostAudioInputStreamSourceClass
}

type _VZHostAudioInputStreamSourceClass struct {
	class objc.Class
}

// An interface definition for the [VZHostAudioInputStreamSource] class.
type IVZHostAudioInputStreamSource interface {
	IVZAudioInputStreamSource
	// properties:
	// methods:
}

// The host audio input stream source that provides audio from the host system’s default input device.
//
// The host input data comes from the same device that uses.


// The host audio input stream source that provides audio from the host system’s default input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZHostAudioInputStreamSource
type VZHostAudioInputStreamSource struct {
	VZAudioInputStreamSource
}

// VZHostAudioInputStreamSourceFrom constructs a [VZHostAudioInputStreamSource] from an unsafe.Pointer.
//
// The host audio input stream source that provides audio from the host system’s default input device.
func VZHostAudioInputStreamSourceFrom(ptr unsafe.Pointer) VZHostAudioInputStreamSource {
	return VZHostAudioInputStreamSource{
		VZAudioInputStreamSource: VZAudioInputStreamSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZHostAudioInputStreamSourceClass) Alloc() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZHostAudioInputStreamSourceClass) New() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZHostAudioInputStreamSource) Init() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZHostAudioInputStreamSource) Autorelease() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioInputStreamSource creates a new VZHostAudioInputStreamSource instance.
func NewVZHostAudioInputStreamSource() VZHostAudioInputStreamSource {
	return getVZHostAudioInputStreamSourceClass().New()
}




