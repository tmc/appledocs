// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetWriterInputPixelBufferAdaptor */


/* debug [class_header]: Header for AVAssetWriterInputPixelBufferAdaptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetWriterInputPixelBufferAdaptor */
// An interface definition for the [AssetWriterInputPixelBufferAdaptor] class.
type IAssetWriterInputPixelBufferAdaptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetWriterInputPixelBufferAdaptor */
	// properties:
	AssetWriterInput() IAVAssetWriterInput
	PixelBufferPool() PixelBufferPoolRef /* not a class type */
	SourcePixelBufferAttributes() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetWriterInputPixelBufferAdaptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetWriterInputPixelBufferAdaptor */
// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputPixelBufferAdaptorClass) Alloc() AssetWriterInputPixelBufferAdaptor {
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetWriterInputPixelBufferAdaptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetWriterInputPixelBufferAdaptor */

// Creates a new pixel buffer adaptor to receive pixel buffers for writing to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/init(assetWriterInput:sourcePixelBufferAttributes:)
func NewAssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAVAssetWriterInput, sourcePixelBufferAttributes foundation.IDictionary) AssetWriterInputPixelBufferAdaptor {
	instance := getAssetWriterInputPixelBufferAdaptorClass().Alloc()
	rv := objc.Send[AssetWriterInputPixelBufferAdaptor](instance.ID, objc.Sel("initWithAssetWriterInput:sourcePixelBufferAttributes:"), input, sourcePixelBufferAttributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetWriterInputPixelBufferAdaptor */

// Returns a new pixel buffer adaptor that appends pixel buffers to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/assetWriterInputPixelBufferAdaptorWithAssetWriterInput:sourcePixelBufferAttributes:
func (ac _AssetWriterInputPixelBufferAdaptorClass) AssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes(input IAVAssetWriterInput, sourcePixelBufferAttributes foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputPixelBufferAdaptorWithAssetWriterInput:sourcePixelBufferAttributes:"), input, sourcePixelBufferAttributes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetWriterInputPixelBufferAdaptorWithAssetWriterInputSourcePixelBufferAttributes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetWriterInputPixelBufferAdaptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetWriterInputPixelBufferAdaptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetWriterInputPixelBufferAdaptor */

// The asset writer input to which the adaptor appends pixel buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/assetWriterInput
func (a_ AssetWriterInputPixelBufferAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}/* debug [instance_properties/getter]: assetWriterInput */


// A pool of pixel buffers to append to the adaptor’s input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/pixelBufferPool
func (a_ AssetWriterInputPixelBufferAdaptor) PixelBufferPool() PixelBufferPoolRef /* not a class type */ {
	rv := objc.Send[PixelBufferPoolRef](a_.ID, objc.Sel("pixelBufferPool"))
	return rv
}/* debug [instance_properties/getter]: pixelBufferPool */


// The attributes of the pixel buffers that the pool contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/sourcePixelBufferAttributes
func (a_ AssetWriterInputPixelBufferAdaptor) SourcePixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("sourcePixelBufferAttributes"))
	return rv
}/* debug [instance_properties/getter]: sourcePixelBufferAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetWriterInputPixelBufferAdaptor */


