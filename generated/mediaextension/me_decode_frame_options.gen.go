// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEDecodeFrameOptions] class.
var (
	MEDecodeFrameOptionsClass     _MEDecodeFrameOptionsClass
	MEDecodeFrameOptionsClassOnce sync.Once
)

func getMEDecodeFrameOptionsClass() _MEDecodeFrameOptionsClass {
	MEDecodeFrameOptionsClassOnce.Do(func() {
		MEDecodeFrameOptionsClass = _MEDecodeFrameOptionsClass{objc.GetClass("MEDecodeFrameOptions")}
	})
	return MEDecodeFrameOptionsClass
}

type _MEDecodeFrameOptionsClass struct {
	class objc.Class
}

// An interface definition for the [MEDecodeFrameOptions] class.
type IMEDecodeFrameOptions interface {
	objectivec.IObject
}

// An object that guides the video decoder operation on a per-frame basis.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions
type MEDecodeFrameOptions struct {
	objectivec.Object
}

// MEDecodeFrameOptionsFrom constructs a [MEDecodeFrameOptions] from an unsafe.Pointer.
//
// An object that guides the video decoder operation on a per-frame basis.
func MEDecodeFrameOptionsFrom(ptr unsafe.Pointer) MEDecodeFrameOptions {
	return MEDecodeFrameOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEDecodeFrameOptionsClass) Alloc() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEDecodeFrameOptionsClass) New() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEDecodeFrameOptions) Init() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEDecodeFrameOptions) Autorelease() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodeFrameOptions creates a new MEDecodeFrameOptions instance.
func NewMEDecodeFrameOptions() MEDecodeFrameOptions {
	return getMEDecodeFrameOptionsClass().New()
}


// A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/doNotOutputFrame
func (m_ MEDecodeFrameOptions) DoNotOutputFrame() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("doNotOutputFrame"))
	return rv
}


// SetDoNotOutputFrame sets the value of the doNotOutputFrame property.
// A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/doNotOutputFrame
func (m_ MEDecodeFrameOptions) SetDoNotOutputFrame(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoNotOutputFrame:"), value)
}

// A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/realTimePlayback
func (m_ MEDecodeFrameOptions) RealTimePlayback() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("realTimePlayback"))
	return rv
}


// SetRealTimePlayback sets the value of the realTimePlayback property.
// A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/realTimePlayback
func (m_ MEDecodeFrameOptions) SetRealTimePlayback(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRealTimePlayback:"), value)
}



