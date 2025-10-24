// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetWriterInputMetadataAdaptor */


/* debug [class_header]: Header for AVAssetWriterInputMetadataAdaptor */
// The class instance for the [AssetWriterInputMetadataAdaptor] class.
var (
	AssetWriterInputMetadataAdaptorClass     _AssetWriterInputMetadataAdaptorClass
	AssetWriterInputMetadataAdaptorClassOnce sync.Once
)

func getAssetWriterInputMetadataAdaptorClass() _AssetWriterInputMetadataAdaptorClass {
	AssetWriterInputMetadataAdaptorClassOnce.Do(func() {
		AssetWriterInputMetadataAdaptorClass = _AssetWriterInputMetadataAdaptorClass{objc.GetClass("AVAssetWriterInputMetadataAdaptor")}
	})
	return AssetWriterInputMetadataAdaptorClass
}

type _AssetWriterInputMetadataAdaptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetWriterInputMetadataAdaptor */
// An interface definition for the [AssetWriterInputMetadataAdaptor] class.
type IAssetWriterInputMetadataAdaptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetWriterInputMetadataAdaptor */
	// properties:
	AssetWriterInput() IAVAssetWriterInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetWriterInputMetadataAdaptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetWriterInputMetadataAdaptor */
// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputMetadataAdaptorClass) Alloc() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetWriterInputMetadataAdaptorClass) New() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputMetadataAdaptor) Init() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputMetadataAdaptor) Autorelease() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputMetadataAdaptor creates a new AssetWriterInputMetadataAdaptor instance.
func NewAssetWriterInputMetadataAdaptor() AssetWriterInputMetadataAdaptor {
	return getAssetWriterInputMetadataAdaptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetWriterInputMetadataAdaptor */
// An object that appends timed metadata groups to an asset writer input.
//
// Use a metadata adaptor to append track-level metadata, packaged as instances of , to an asset writer input.


// An object that appends timed metadata groups to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor
type AssetWriterInputMetadataAdaptor struct {
	objectivec.Object
}

// AssetWriterInputMetadataAdaptorFrom constructs a [AssetWriterInputMetadataAdaptor] from an unsafe.Pointer.
//
// An object that appends timed metadata groups to an asset writer input.
func AssetWriterInputMetadataAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetWriterInputMetadataAdaptor */

// Creates a metadata group adaptor to append timed metadata groups to write to an output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/init(assetWriterInput:)
func NewAssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAVAssetWriterInput) AssetWriterInputMetadataAdaptor {
	instance := getAssetWriterInputMetadataAdaptorClass().Alloc()
	rv := objc.Send[AssetWriterInputMetadataAdaptor](instance.ID, objc.Sel("initWithAssetWriterInput:"), input)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetWriterInputMetadataAdaptorWithAssetWriterInput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetWriterInputMetadataAdaptor */

// Returns a new metadata adaptor to append timed metadata groups to write to an output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/assetWriterInputMetadataAdaptorWithAssetWriterInput:
func (ac _AssetWriterInputMetadataAdaptorClass) AssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAVAssetWriterInput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputMetadataAdaptorWithAssetWriterInput:"), input)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetWriterInputMetadataAdaptorWithAssetWriterInput) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetWriterInputMetadataAdaptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetWriterInputMetadataAdaptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetWriterInputMetadataAdaptor */

// The input for the metadata adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/assetWriterInput
func (a_ AssetWriterInputMetadataAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}/* debug [instance_properties/getter]: assetWriterInput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetWriterInputMetadataAdaptor */


