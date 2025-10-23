// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZAudioInputStreamSource] class.
var (
	VZAudioInputStreamSourceClass     _VZAudioInputStreamSourceClass
	VZAudioInputStreamSourceClassOnce sync.Once
)

func getVZAudioInputStreamSourceClass() _VZAudioInputStreamSourceClass {
	VZAudioInputStreamSourceClassOnce.Do(func() {
		VZAudioInputStreamSourceClass = _VZAudioInputStreamSourceClass{objc.GetClass("VZAudioInputStreamSource")}
	})
	return VZAudioInputStreamSourceClass
}

type _VZAudioInputStreamSourceClass struct {
	class objc.Class
}

// An interface definition for the [VZAudioInputStreamSource] class.
type IVZAudioInputStreamSource interface {
	objectivec.IObject
}

// The base class for an audio input stream source.
//
// An audio input stream source defines how th guest produces audio input data on the host system. Don’t instantiate directly, use one of its subclasses such as instead.


// The base class for an audio input stream source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type VZAudioInputStreamSource struct {
	objectivec.Object
}

// VZAudioInputStreamSourceFrom constructs a [VZAudioInputStreamSource] from an unsafe.Pointer.
//
// The base class for an audio input stream source.
func VZAudioInputStreamSourceFrom(ptr unsafe.Pointer) VZAudioInputStreamSource {
	return VZAudioInputStreamSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZAudioInputStreamSourceClass) Alloc() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZAudioInputStreamSourceClass) New() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioInputStreamSource) Init() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioInputStreamSource) Autorelease() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioInputStreamSource creates a new VZAudioInputStreamSource instance.
func NewVZAudioInputStreamSource() VZAudioInputStreamSource {
	return getVZAudioInputStreamSourceClass().New()
}




