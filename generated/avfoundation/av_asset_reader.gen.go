// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReader */


/* debug [class_header]: Header for AVAssetReader */
// The class instance for the [AssetReader] class.
var (
	AssetReaderClass     _AssetReaderClass
	AssetReaderClassOnce sync.Once
)

func getAssetReaderClass() _AssetReaderClass {
	AssetReaderClassOnce.Do(func() {
		AssetReaderClass = _AssetReaderClass{objc.GetClass("AVAssetReader")}
	})
	return AssetReaderClass
}

type _AssetReaderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReader */
// An interface definition for the [AssetReader] class.
type IAssetReader interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetReader */
	// properties:
	Asset() IAVAsset
	Error() Error
	Outputs() []AssetReaderOutput
	Status() AssetReaderStatus
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReader */
	// methods:
	CanAddOutput(output IAVAssetReaderOutput) bool
	CancelReading()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReader */
// Alloc allocates a new instance without initialization.
func (ac _AssetReaderClass) Alloc() AssetReader {
	rv := objc.Send[AssetReader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderClass) New() AssetReader {
	rv := objc.Send[AssetReader](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReader) Init() AssetReader {
	rv := objc.Send[AssetReader](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReader) Autorelease() AssetReader {
	rv := objc.Send[AssetReader](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReader creates a new AssetReader instance.
func NewAssetReader() AssetReader {
	return getAssetReaderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReader */
// An object that reads media data from an asset.
//
// Use an asset reader to read media data from instances of . The assets you read may represent file-based media like QuickTime movies or MPEG-4 files, or media that you compose from multiple sources using .


// An object that reads media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader
type AssetReader struct {
	objectivec.Object
}

// AssetReaderFrom constructs a [AssetReader] from an unsafe.Pointer.
//
// An object that reads media data from an asset.
func AssetReaderFrom(ptr unsafe.Pointer) AssetReader {
	return AssetReader{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReader */

// Creates an object to read media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/init(asset:)
func NewAssetReaderWithAssetError(asset IAVAsset, outError objectivec.IObject) AssetReader {
	instance := getAssetReaderClass().Alloc()
	rv := objc.Send[AssetReader](instance.ID, objc.Sel("initWithAsset:error:"), asset, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetReaderWithAssetError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReader */

// Returns a new object to read media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/assetReaderWithAsset:error:
func (ac _AssetReaderClass) AssetReaderWithAssetError(asset IAVAsset, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderWithAsset:error:"), asset, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetReaderWithAssetError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReader */

// Determines whether you can add the output to the asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/canAdd(_:)
func (a_ AssetReader) CanAddOutput(output IAVAssetReaderOutput) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}/* debug [instance_methods/method]: CanAddOutput */


// Cancels any background work and stops the reader’s outputs from reading more samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/cancelReading()
func (a_ AssetReader) CancelReading() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelReading"))
}/* debug [instance_methods/method]: CancelReading */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReader */

// The asset from which to read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/asset
func (a_ AssetReader) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */


// An error that describes the reason for a failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/error
func (a_ AssetReader) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The outputs from which you read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/outputs
func (a_ AssetReader) Outputs() []AssetReaderOutput {
	rv := objc.Send[[]AssetReaderOutput](a_.ID, objc.Sel("outputs"))
	return rv
}/* debug [instance_properties/getter]: outputs */


// The status of reading sample buffers from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/status-swift.property
func (a_ AssetReader) Status() AssetReaderStatus {
	rv := objc.Send[AssetReaderStatus](a_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/timeRange
func (a_ AssetReader) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/timeRange
func (a_ AssetReader) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}/* debug [instance_properties/setter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReader */


