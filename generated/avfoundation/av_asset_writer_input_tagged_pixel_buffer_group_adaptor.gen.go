// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetWriterInputTaggedPixelBufferGroupAdaptor] class.
var (
	AssetWriterInputTaggedPixelBufferGroupAdaptorClass     _AssetWriterInputTaggedPixelBufferGroupAdaptorClass
	AssetWriterInputTaggedPixelBufferGroupAdaptorClassOnce sync.Once
)

func getAssetWriterInputTaggedPixelBufferGroupAdaptorClass() _AssetWriterInputTaggedPixelBufferGroupAdaptorClass {
	AssetWriterInputTaggedPixelBufferGroupAdaptorClassOnce.Do(func() {
		AssetWriterInputTaggedPixelBufferGroupAdaptorClass = _AssetWriterInputTaggedPixelBufferGroupAdaptorClass{objc.GetClass("AVAssetWriterInputTaggedPixelBufferGroupAdaptor")}
	})
	return AssetWriterInputTaggedPixelBufferGroupAdaptorClass
}

type _AssetWriterInputTaggedPixelBufferGroupAdaptorClass struct {
	class objc.Class
}





// An interface definition for the [AssetWriterInputTaggedPixelBufferGroupAdaptor] class.
type IAssetWriterInputTaggedPixelBufferGroupAdaptor interface {
	objectivec.IObject
	

	// properties:
	AssetWriterInput() IAVAssetWriterInput
	PixelBufferPool() PixelBufferPoolRef /* not a class type */
	SourcePixelBufferAttributes() foundation.IDictionary


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputTaggedPixelBufferGroupAdaptorClass) Alloc() AssetWriterInputTaggedPixelBufferGroupAdaptor {
	rv := objc.Send[AssetWriterInputTaggedPixelBufferGroupAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetWriterInputTaggedPixelBufferGroupAdaptorClass) New() AssetWriterInputTaggedPixelBufferGroupAdaptor {
	rv := objc.Send[AssetWriterInputTaggedPixelBufferGroupAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputTaggedPixelBufferGroupAdaptor) Init() AssetWriterInputTaggedPixelBufferGroupAdaptor {
	rv := objc.Send[AssetWriterInputTaggedPixelBufferGroupAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputTaggedPixelBufferGroupAdaptor) Autorelease() AssetWriterInputTaggedPixelBufferGroupAdaptor {
	rv := objc.Send[AssetWriterInputTaggedPixelBufferGroupAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputTaggedPixelBufferGroupAdaptor creates a new AssetWriterInputTaggedPixelBufferGroupAdaptor instance.
func NewAssetWriterInputTaggedPixelBufferGroupAdaptor() AssetWriterInputTaggedPixelBufferGroupAdaptor {
	return getAssetWriterInputTaggedPixelBufferGroupAdaptorClass().New()
}





// An object that appends tagged buffer groups to an asset writer input.
//
// This class provides a to use for allocating the pixel buffers of tagged buffer groups to write to the output file. Using the provided pixel buffer pool for buffer allocation is typically more efficient than appending pixel buffers allocated using a separate pool.


// An object that appends tagged buffer groups to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor
type AssetWriterInputTaggedPixelBufferGroupAdaptor struct {
	objectivec.Object
}

// AssetWriterInputTaggedPixelBufferGroupAdaptorFrom constructs a [AssetWriterInputTaggedPixelBufferGroupAdaptor] from an unsafe.Pointer.
//
// An object that appends tagged buffer groups to an asset writer input.
func AssetWriterInputTaggedPixelBufferGroupAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputTaggedPixelBufferGroupAdaptor {
	return AssetWriterInputTaggedPixelBufferGroupAdaptor{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that appends tagged buffer groups to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor/init(assetWriterInput:sourcePixelBufferAttributes:)
func NewAssetWriterInputTaggedPixelBufferGroupAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAVAssetWriterInput, sourcePixelBufferAttributes foundation.IDictionary) AssetWriterInputTaggedPixelBufferGroupAdaptor {
	instance := getAssetWriterInputTaggedPixelBufferGroupAdaptorClass().Alloc()
	rv := objc.Send[AssetWriterInputTaggedPixelBufferGroupAdaptor](instance.ID, objc.Sel("initWithAssetWriterInput:sourcePixelBufferAttributes:"), input, sourcePixelBufferAttributes)
	rv.Autorelease()
	return rv
}







// Returns a new object that appends tagged buffer groups to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor/assetWriterInputTaggedPixelBufferGroupAdaptorWithAssetWriterInput:sourcePixelBufferAttributes:
func (ac _AssetWriterInputTaggedPixelBufferGroupAdaptorClass) AssetWriterInputTaggedPixelBufferGroupAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAVAssetWriterInput, sourcePixelBufferAttributes foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputTaggedPixelBufferGroupAdaptorWithAssetWriterInput:sourcePixelBufferAttributes:"), input, sourcePixelBufferAttributes)
	return rv
}

















// The asset writer input to which the adaptor appends tagged buffer groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor/assetWriterInput
func (a_ AssetWriterInputTaggedPixelBufferGroupAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}


// A pixel buffer pool that vends and efficiently recycles the pixel buffers of tagged buffer groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor/pixelBufferPool
func (a_ AssetWriterInputTaggedPixelBufferGroupAdaptor) PixelBufferPool() PixelBufferPoolRef /* not a class type */ {
	rv := objc.Send[PixelBufferPoolRef](a_.ID, objc.Sel("pixelBufferPool"))
	return rv
}


// The attributes of buffers that the adaptor’s pixel buffer pool vends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputTaggedPixelBufferGroupAdaptor/sourcePixelBufferAttributes
func (a_ AssetWriterInputTaggedPixelBufferGroupAdaptor) SourcePixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("sourcePixelBufferAttributes"))
	return rv
}







