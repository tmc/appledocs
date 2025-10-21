// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MEVideoDecoderPixelBufferManager] class.
var (
	MEVideoDecoderPixelBufferManagerClass     _MEVideoDecoderPixelBufferManagerClass
	MEVideoDecoderPixelBufferManagerClassOnce sync.Once
)

func getMEVideoDecoderPixelBufferManagerClass() _MEVideoDecoderPixelBufferManagerClass {
	MEVideoDecoderPixelBufferManagerClassOnce.Do(func() {
		MEVideoDecoderPixelBufferManagerClass = _MEVideoDecoderPixelBufferManagerClass{objc.GetClass("MEVideoDecoderPixelBufferManager")}
	})
	return MEVideoDecoderPixelBufferManagerClass
}

type _MEVideoDecoderPixelBufferManagerClass struct {
	class objc.Class
}

// An interface definition for the [MEVideoDecoderPixelBufferManager] class.
type IMEVideoDecoderPixelBufferManager interface {
	objectivec.IObject
	RegisterCustomPixelFormat(customPixelFormat unsafe.Pointer)
}

// Describes pixel buffer requirements and creates new pixel buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager
type MEVideoDecoderPixelBufferManager struct {
	objectivec.Object
}

// MEVideoDecoderPixelBufferManagerFrom constructs a [MEVideoDecoderPixelBufferManager] from an unsafe.Pointer.
//
// Describes pixel buffer requirements and creates new pixel buffers.
func MEVideoDecoderPixelBufferManagerFrom(ptr unsafe.Pointer) MEVideoDecoderPixelBufferManager {
	return MEVideoDecoderPixelBufferManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEVideoDecoderPixelBufferManagerClass) Alloc() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEVideoDecoderPixelBufferManagerClass) New() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEVideoDecoderPixelBufferManager) Init() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEVideoDecoderPixelBufferManager) Autorelease() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEVideoDecoderPixelBufferManager creates a new MEVideoDecoderPixelBufferManager instance.
func NewMEVideoDecoderPixelBufferManager() MEVideoDecoderPixelBufferManager {
	return getMEVideoDecoderPixelBufferManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/registerCustomPixelFormat(_:)
func (m_ MEVideoDecoderPixelBufferManager) RegisterCustomPixelFormat(customPixelFormat unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerCustomPixelFormat:"), customPixelFormat)
}

// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/pixelBufferAttributes
func (m_ MEVideoDecoderPixelBufferManager) PixelBufferAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// SetPixelBufferAttributes sets the value of the pixelBufferAttributes property.
// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/pixelBufferAttributes
func (m_ MEVideoDecoderPixelBufferManager) SetPixelBufferAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}




