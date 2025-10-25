// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReaderOutputCaptionAdaptor */


/* debug [class_header]: Header for AVAssetReaderOutputCaptionAdaptor */
// The class instance for the [AssetReaderOutputCaptionAdaptor] class.
var (
	AssetReaderOutputCaptionAdaptorClass     _AssetReaderOutputCaptionAdaptorClass
	AssetReaderOutputCaptionAdaptorClassOnce sync.Once
)

func getAssetReaderOutputCaptionAdaptorClass() _AssetReaderOutputCaptionAdaptorClass {
	AssetReaderOutputCaptionAdaptorClassOnce.Do(func() {
		AssetReaderOutputCaptionAdaptorClass = _AssetReaderOutputCaptionAdaptorClass{objc.GetClass("AVAssetReaderOutputCaptionAdaptor")}
	})
	return AssetReaderOutputCaptionAdaptorClass
}

type _AssetReaderOutputCaptionAdaptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReaderOutputCaptionAdaptor */
// An interface definition for the [AssetReaderOutputCaptionAdaptor] class.
type IAssetReaderOutputCaptionAdaptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetReaderOutputCaptionAdaptor */
	// properties:
	AssetReaderTrackOutput() IAVAssetReaderTrackOutput
	ValidationDelegate() unsafe.Pointer
	SetValidationDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReaderOutputCaptionAdaptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReaderOutputCaptionAdaptor */
// Alloc allocates a new instance without initialization.
func (ac _AssetReaderOutputCaptionAdaptorClass) Alloc() AssetReaderOutputCaptionAdaptor {
	rv := objc.Send[AssetReaderOutputCaptionAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderOutputCaptionAdaptorClass) New() AssetReaderOutputCaptionAdaptor {
	rv := objc.Send[AssetReaderOutputCaptionAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderOutputCaptionAdaptor) Init() AssetReaderOutputCaptionAdaptor {
	rv := objc.Send[AssetReaderOutputCaptionAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderOutputCaptionAdaptor) Autorelease() AssetReaderOutputCaptionAdaptor {
	rv := objc.Send[AssetReaderOutputCaptionAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderOutputCaptionAdaptor creates a new AssetReaderOutputCaptionAdaptor instance.
func NewAssetReaderOutputCaptionAdaptor() AssetReaderOutputCaptionAdaptor {
	return getAssetReaderOutputCaptionAdaptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReaderOutputCaptionAdaptor */
// An object that reads caption group objects from an asset track that contains timed text.


// An object that reads caption group objects from an asset track that contains timed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor
type AssetReaderOutputCaptionAdaptor struct {
	objectivec.Object
}

// AssetReaderOutputCaptionAdaptorFrom constructs a [AssetReaderOutputCaptionAdaptor] from an unsafe.Pointer.
//
// An object that reads caption group objects from an asset track that contains timed text.
func AssetReaderOutputCaptionAdaptorFrom(ptr unsafe.Pointer) AssetReaderOutputCaptionAdaptor {
	return AssetReaderOutputCaptionAdaptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReaderOutputCaptionAdaptor */

// Creates a caption adaptor that reads from a track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor/init(assetReaderTrackOutput:)
func NewAssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput IAVAssetReaderTrackOutput) AssetReaderOutputCaptionAdaptor {
	instance := getAssetReaderOutputCaptionAdaptorClass().Alloc()
	rv := objc.Send[AssetReaderOutputCaptionAdaptor](instance.ID, objc.Sel("initWithAssetReaderTrackOutput:"), trackOutput)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReaderOutputCaptionAdaptor */

// A class method that creates a caption adaptor that reads from a track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor/assetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput:
func (ac _AssetReaderOutputCaptionAdaptorClass) AssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput IAVAssetReaderTrackOutput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput:"), trackOutput)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReaderOutputCaptionAdaptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReaderOutputCaptionAdaptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReaderOutputCaptionAdaptor */

// The associated asset reader track output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor/assetReaderTrackOutput
func (a_ AssetReaderOutputCaptionAdaptor) AssetReaderTrackOutput() IAVAssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("assetReaderTrackOutput"))
	return rv
}/* debug [instance_properties/getter]: assetReaderTrackOutput */


// A delegate object that handles callbacks to the caption adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor/validationDelegate
func (a_ AssetReaderOutputCaptionAdaptor) ValidationDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("validationDelegate"))
	return rv
}/* debug [instance_properties/getter]: validationDelegate */


// A delegate object that handles callbacks to the caption adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputCaptionAdaptor/validationDelegate
func (a_ AssetReaderOutputCaptionAdaptor) SetValidationDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValidationDelegate:"), value)
}/* debug [instance_properties/setter]: validationDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReaderOutputCaptionAdaptor */


