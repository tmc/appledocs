// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEVideoDecoderPixelBufferManager */


/* debug [class_header]: Header for MEVideoDecoderPixelBufferManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEVideoDecoderPixelBufferManager */
// An interface definition for the [MEVideoDecoderPixelBufferManager] class.
type IMEVideoDecoderPixelBufferManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEVideoDecoderPixelBufferManager */
	// properties:
	PixelBufferAttributes() foundation.IDictionary
	SetPixelBufferAttributes(value foundation.IDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEVideoDecoderPixelBufferManager */
	// methods:
	CreatePixelBufferAndReturnError(error_ unsafe.Pointer) PixelBufferRef /* not a class type */
	RegisterCustomPixelFormat(customPixelFormat foundation.IDictionary)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEVideoDecoderPixelBufferManager */
// Alloc allocates a new instance without initialization.
func (mc _MEVideoDecoderPixelBufferManagerClass) Alloc() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEVideoDecoderPixelBufferManager */
// Describes pixel buffer requirements and creates new pixel buffers.
//
// Contains the interfaces that the uses for two tasks. First, to declare its set of requirements for output objects in the form of a dictionary. Second, to create pixel buffers that match decoder output requirements but also satisfy and client requirements.


// Describes pixel buffer requirements and creates new pixel buffers.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEVideoDecoderPixelBufferManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEVideoDecoderPixelBufferManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEVideoDecoderPixelBufferManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEVideoDecoderPixelBufferManager */

// Generates a pixel buffer using the session’s pixel buffer pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/makePixelBuffer()
func (m_ MEVideoDecoderPixelBufferManager) CreatePixelBufferAndReturnError(error_ unsafe.Pointer) PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](m_.ID, objc.Sel("createPixelBufferAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: CreatePixelBufferAndReturnError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/registerCustomPixelFormat(_:)
func (m_ MEVideoDecoderPixelBufferManager) RegisterCustomPixelFormat(customPixelFormat foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerCustomPixelFormat:"), customPixelFormat)
}/* debug [instance_methods/method]: RegisterCustomPixelFormat */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEVideoDecoderPixelBufferManager */

// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/pixelBufferAttributes
func (m_ MEVideoDecoderPixelBufferManager) PixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}/* debug [instance_properties/getter]: pixelBufferAttributes */


// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/pixelBufferAttributes
func (m_ MEVideoDecoderPixelBufferManager) SetPixelBufferAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}/* debug [instance_properties/setter]: pixelBufferAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEVideoDecoderPixelBufferManager */





