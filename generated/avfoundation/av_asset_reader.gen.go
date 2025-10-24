// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [AssetReader] class.
type IAssetReader interface {
	objectivec.IObject
	

	// properties:
	Asset() IAVAsset
	Error() Error
	Outputs() []AssetReaderOutput
	Status() AssetReaderStatus
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)


	

	// methods:
	CanAddOutput(output IAVAssetReaderOutput) bool
	CancelReading()


}





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






// Creates an object to read media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/init(asset:)
func NewAssetReaderWithAssetError(asset IAVAsset, outError objectivec.IObject) AssetReader {
	instance := getAssetReaderClass().Alloc()
	rv := objc.Send[AssetReader](instance.ID, objc.Sel("initWithAsset:error:"), asset, outError)
	rv.Autorelease()
	return rv
}







// Returns a new object to read media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/assetReaderWithAsset:error:
func (ac _AssetReaderClass) AssetReaderWithAssetError(asset IAVAsset, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderWithAsset:error:"), asset, outError)
	return rv
}












// Determines whether you can add the output to the asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/canAdd(_:)
func (a_ AssetReader) CanAddOutput(output IAVAssetReaderOutput) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}


// Cancels any background work and stops the reader’s outputs from reading more samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/cancelReading()
func (a_ AssetReader) CancelReading() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelReading"))
}







// The asset from which to read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/asset
func (a_ AssetReader) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// An error that describes the reason for a failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/error
func (a_ AssetReader) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// The outputs from which you read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/outputs
func (a_ AssetReader) Outputs() []AssetReaderOutput {
	rv := objc.Send[[]AssetReaderOutput](a_.ID, objc.Sel("outputs"))
	return rv
}


// The status of reading sample buffers from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/status-swift.property
func (a_ AssetReader) Status() AssetReaderStatus {
	rv := objc.Send[AssetReaderStatus](a_.ID, objc.Sel("status"))
	return rv
}


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/timeRange
func (a_ AssetReader) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/timeRange
func (a_ AssetReader) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}







