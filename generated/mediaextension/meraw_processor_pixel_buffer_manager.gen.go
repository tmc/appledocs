// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	PixelBufferAttributes() foundation.IDictionary
	SetPixelBufferAttributes(value foundation.IDictionary)
	// methods:
	CreatePixelBufferAndReturnError(error_ unsafe.Pointer) PixelBufferRef /* not a class type */
}

// Describes pixel buffer requirements and creates new pixel buffers.
//
// It contains the interfaces that the uses for two tasks. First, to declare its set of requirements for output in the form of a dictionary. Second, create pixel buffers that match processor output requirements and satisfy Video Toolbox and client requirements.


// Describes pixel buffer requirements and creates new pixel buffers.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/makePixelBuffer()
func (m_ MERAWProcessorPixelBufferManager) CreatePixelBufferAndReturnError(error_ unsafe.Pointer) PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](m_.ID, objc.Sel("createPixelBufferAndReturnError:"), error_)
	return rv
}


// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/pixelBufferAttributes-2cki6
func (m_ MERAWProcessorPixelBufferManager) PixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/pixelBufferAttributes-2cki6
func (m_ MERAWProcessorPixelBufferManager) SetPixelBufferAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}



