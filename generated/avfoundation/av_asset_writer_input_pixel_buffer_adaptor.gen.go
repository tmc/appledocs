// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetWriterInputPixelBufferAdaptor] class.
var (
	AssetWriterInputPixelBufferAdaptorClass     _AssetWriterInputPixelBufferAdaptorClass
	AssetWriterInputPixelBufferAdaptorClassOnce sync.Once
)

func getAssetWriterInputPixelBufferAdaptorClass() _AssetWriterInputPixelBufferAdaptorClass {
	AssetWriterInputPixelBufferAdaptorClassOnce.Do(func() {
		AssetWriterInputPixelBufferAdaptorClass = _AssetWriterInputPixelBufferAdaptorClass{objc.GetClass("AVAssetWriterInputPixelBufferAdaptor")}
	})
	return AssetWriterInputPixelBufferAdaptorClass
}

type _AssetWriterInputPixelBufferAdaptorClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriterInputPixelBufferAdaptor] class.
type IAssetWriterInputPixelBufferAdaptor interface {
	objectivec.IObject
	// properties:
	AssetWriterInput() IAVAssetWriterInput
	SetAssetWriterInput(value IAVAssetWriterInput)
	PixelBufferPool() PixelBufferPool /* not a class type */
	SetPixelBufferPool(value PixelBufferPool /* not a class type */)
	SourcePixelBufferAttributes() unsafe.Pointer
	SetSourcePixelBufferAttributes(value unsafe.Pointer)
	// methods:
}

// An object that appends video samples to an asset writer input.
//
// A pixel buffer adaptor provides a pixel buffer pool that you use to allocate pixel buffers to the output file. Using the provided pool for buffer allocation is typically more efficient than managing your own pool.


// An object that appends video samples to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor
type AssetWriterInputPixelBufferAdaptor struct {
	objectivec.Object
}

// AssetWriterInputPixelBufferAdaptorFrom constructs a [AssetWriterInputPixelBufferAdaptor] from an unsafe.Pointer.
//
// An object that appends video samples to an asset writer input.
func AssetWriterInputPixelBufferAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputPixelBufferAdaptor {
	return AssetWriterInputPixelBufferAdaptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputPixelBufferAdaptorClass) Alloc() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterInputPixelBufferAdaptorClass) New() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputPixelBufferAdaptor) Init() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputPixelBufferAdaptor) Autorelease() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputPixelBufferAdaptor creates a new AssetWriterInputPixelBufferAdaptor instance.
func NewAssetWriterInputPixelBufferAdaptor() AssetWriterInputPixelBufferAdaptor {
	return getAssetWriterInputPixelBufferAdaptorClass().New()
}



// The asset writer input to which the adaptor appends pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/assetwriterinput
func (a_ AssetWriterInputPixelBufferAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}


// The asset writer input to which the adaptor appends pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/assetwriterinput
func (a_ AssetWriterInputPixelBufferAdaptor) SetAssetWriterInput(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAssetWriterInput:"), value)
}


// A pool of pixel buffers to append to the adaptor’s input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/pixelbufferpool
func (a_ AssetWriterInputPixelBufferAdaptor) PixelBufferPool() PixelBufferPool /* not a class type */ {
	rv := objc.Send[PixelBufferPool](a_.ID, objc.Sel("pixelBufferPool"))
	return rv
}


// A pool of pixel buffers to append to the adaptor’s input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/pixelbufferpool
func (a_ AssetWriterInputPixelBufferAdaptor) SetPixelBufferPool(value PixelBufferPool /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPixelBufferPool:"), value)
}


// The attributes of the pixel buffers that the pool contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/sourcepixelbufferattributes
func (a_ AssetWriterInputPixelBufferAdaptor) SourcePixelBufferAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sourcePixelBufferAttributes"))
	return rv
}


// The attributes of the pixel buffers that the pool contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputpixelbufferadaptor/sourcepixelbufferattributes
func (a_ AssetWriterInputPixelBufferAdaptor) SetSourcePixelBufferAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSourcePixelBufferAttributes:"), value)
}



