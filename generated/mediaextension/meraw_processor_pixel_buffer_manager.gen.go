// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MERAWProcessorPixelBufferManager] class.
var (
	MERAWProcessorPixelBufferManagerClass     _MERAWProcessorPixelBufferManagerClass
	MERAWProcessorPixelBufferManagerClassOnce sync.Once
)

func getMERAWProcessorPixelBufferManagerClass() _MERAWProcessorPixelBufferManagerClass {
	MERAWProcessorPixelBufferManagerClassOnce.Do(func() {
		MERAWProcessorPixelBufferManagerClass = _MERAWProcessorPixelBufferManagerClass{objc.GetClass("MERAWProcessorPixelBufferManager")}
	})
	return MERAWProcessorPixelBufferManagerClass
}

type _MERAWProcessorPixelBufferManagerClass struct {
	class objc.Class
}

// An interface definition for the [MERAWProcessorPixelBufferManager] class.
type IMERAWProcessorPixelBufferManager interface {
	objectivec.IObject
	CreatePixelBufferAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
}

// Describes pixel buffer requirements and creates new pixel buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager
type MERAWProcessorPixelBufferManager struct {
	objectivec.Object
}

// MERAWProcessorPixelBufferManagerFrom constructs a [MERAWProcessorPixelBufferManager] from an unsafe.Pointer.
//
// Describes pixel buffer requirements and creates new pixel buffers.
func MERAWProcessorPixelBufferManagerFrom(ptr unsafe.Pointer) MERAWProcessorPixelBufferManager {
	return MERAWProcessorPixelBufferManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessorPixelBufferManagerClass) Alloc() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MERAWProcessorPixelBufferManagerClass) New() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessorPixelBufferManager) Init() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessorPixelBufferManager) Autorelease() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessorPixelBufferManager creates a new MERAWProcessorPixelBufferManager instance.
func NewMERAWProcessorPixelBufferManager() MERAWProcessorPixelBufferManager {
	return getMERAWProcessorPixelBufferManagerClass().New()
}


// Generates a pixel buffer using the session’s pixel buffer pool.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/makePixelBuffer()
func (m_ MERAWProcessorPixelBufferManager) CreatePixelBufferAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("createPixelBufferAndReturnError:"), error_)
	return rv
}

// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/pixelBufferAttributes-2cki6
func (m_ MERAWProcessorPixelBufferManager) PixelBufferAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// SetPixelBufferAttributes sets the value of the pixelBufferAttributes property.
// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/pixelBufferAttributes-2cki6
func (m_ MERAWProcessorPixelBufferManager) SetPixelBufferAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}


