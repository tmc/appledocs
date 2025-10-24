// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetReaderOutputMetadataAdaptor] class.
var (
	AssetReaderOutputMetadataAdaptorClass     _AssetReaderOutputMetadataAdaptorClass
	AssetReaderOutputMetadataAdaptorClassOnce sync.Once
)

func getAssetReaderOutputMetadataAdaptorClass() _AssetReaderOutputMetadataAdaptorClass {
	AssetReaderOutputMetadataAdaptorClassOnce.Do(func() {
		AssetReaderOutputMetadataAdaptorClass = _AssetReaderOutputMetadataAdaptorClass{objc.GetClass("AVAssetReaderOutputMetadataAdaptor")}
	})
	return AssetReaderOutputMetadataAdaptorClass
}

type _AssetReaderOutputMetadataAdaptorClass struct {
	class objc.Class
}





// An interface definition for the [AssetReaderOutputMetadataAdaptor] class.
type IAssetReaderOutputMetadataAdaptor interface {
	objectivec.IObject
	

	// properties:
	AssetReaderTrackOutput() IAVAssetReaderTrackOutput


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetReaderOutputMetadataAdaptorClass) Alloc() AssetReaderOutputMetadataAdaptor {
	rv := objc.Send[AssetReaderOutputMetadataAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderOutputMetadataAdaptorClass) New() AssetReaderOutputMetadataAdaptor {
	rv := objc.Send[AssetReaderOutputMetadataAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderOutputMetadataAdaptor) Init() AssetReaderOutputMetadataAdaptor {
	rv := objc.Send[AssetReaderOutputMetadataAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderOutputMetadataAdaptor) Autorelease() AssetReaderOutputMetadataAdaptor {
	rv := objc.Send[AssetReaderOutputMetadataAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderOutputMetadataAdaptor creates a new AssetReaderOutputMetadataAdaptor instance.
func NewAssetReaderOutputMetadataAdaptor() AssetReaderOutputMetadataAdaptor {
	return getAssetReaderOutputMetadataAdaptorClass().New()
}





// An object that creates timed metadata group objects for an asset track.


// An object that creates timed metadata group objects for an asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputMetadataAdaptor
type AssetReaderOutputMetadataAdaptor struct {
	objectivec.Object
}

// AssetReaderOutputMetadataAdaptorFrom constructs a [AssetReaderOutputMetadataAdaptor] from an unsafe.Pointer.
//
// An object that creates timed metadata group objects for an asset track.
func AssetReaderOutputMetadataAdaptorFrom(ptr unsafe.Pointer) AssetReaderOutputMetadataAdaptor {
	return AssetReaderOutputMetadataAdaptor{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that reads timed metadata groups from an asset reader output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputMetadataAdaptor/init(assetReaderTrackOutput:)
func NewAssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput IAVAssetReaderTrackOutput) AssetReaderOutputMetadataAdaptor {
	instance := getAssetReaderOutputMetadataAdaptorClass().Alloc()
	rv := objc.Send[AssetReaderOutputMetadataAdaptor](instance.ID, objc.Sel("initWithAssetReaderTrackOutput:"), trackOutput)
	rv.Autorelease()
	return rv
}







// Returns a new object that reads timed metadata groups from an asset reader output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputMetadataAdaptor/assetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput:
func (ac _AssetReaderOutputMetadataAdaptorClass) AssetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput(trackOutput IAVAssetReaderTrackOutput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderOutputMetadataAdaptorWithAssetReaderTrackOutput:"), trackOutput)
	return rv
}

















// The asset reader track output that provides the timed metadata groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutputMetadataAdaptor/assetReaderTrackOutput
func (a_ AssetReaderOutputMetadataAdaptor) AssetReaderTrackOutput() IAVAssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("assetReaderTrackOutput"))
	return rv
}







