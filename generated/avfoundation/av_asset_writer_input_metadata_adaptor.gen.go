// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [AssetWriterInputMetadataAdaptor] class.
type IAssetWriterInputMetadataAdaptor interface {
	objectivec.IObject
	

	// properties:
	AssetWriterInput() IAVAssetWriterInput


	

	// methods:


}





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






// Creates a metadata group adaptor to append timed metadata groups to write to an output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/init(assetWriterInput:)
func NewAssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAVAssetWriterInput) AssetWriterInputMetadataAdaptor {
	instance := getAssetWriterInputMetadataAdaptorClass().Alloc()
	rv := objc.Send[AssetWriterInputMetadataAdaptor](instance.ID, objc.Sel("initWithAssetWriterInput:"), input)
	rv.Autorelease()
	return rv
}







// Returns a new metadata adaptor to append timed metadata groups to write to an output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/assetWriterInputMetadataAdaptorWithAssetWriterInput:
func (ac _AssetWriterInputMetadataAdaptorClass) AssetWriterInputMetadataAdaptorWithAssetWriterInput(input IAVAssetWriterInput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputMetadataAdaptorWithAssetWriterInput:"), input)
	return rv
}

















// The input for the metadata adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor/assetWriterInput
func (a_ AssetWriterInputMetadataAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}







